package controllers

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/J-Nokwal/Guess_The_Logo/Server/models"
	"github.com/J-Nokwal/Guess_The_Logo/pb"
)

func (server *Server) StartGame(stream pb.Game_StartGameServer) error {
	userAction, err := stream.Recv()
	if err != nil {
		fmt.Println("Error in Starting Game", err)
		return fmt.Errorf("Error In Starting Game")
	}
	action := userAction.GetAction()
	var level pb.Level
	switch action.(type) {
	case *pb.UserAction_GameRequest:
		level = userAction.GetGameRequest().GetLevel()
		fmt.Println("Level Is Set To:", level)
	case *pb.UserAction_QuestionAnswer:
		return fmt.Errorf("Error in Starting Game, Cant Send Answer Before Game Start")
	}
	var timePerQuestion int32 = 2
	stream.Send(&pb.GameStatus{
		GameId: 1,
		Status: &pb.GameStatus_StartStatus{
			StartStatus: &pb.StartStatus{
				NumberOfQuestions: 10,
				TimePerQueston:    timePerQuestion,
			},
		},
	})
	var nextQuestionChannel = make(chan int)
	var closeChan = make(chan error)
	result := &pb.ResultStatus{}
	go sendQuestionStream(stream, nextQuestionChannel, closeChan, result, timePerQuestion)
	go recieveAnswerStream(stream, nextQuestionChannel)

	select {
	case err := <-closeChan:
		if err != nil {
			return err
		}
		if !(len(result.Ans) == len(result.ScoreOfEachQuestion) && len(result.ScoreOfEachQuestion) == len(result.Question)) {
			return fmt.Errorf("len of result.Ans!= result.ScoreOfEachQuestion 1result.ScoreOfEachQuestion")
		}
		stream.Send(
			&pb.GameStatus{
				GameId: 1,
				Status: &pb.GameStatus_ResultStatus{
					ResultStatus: result,
				},
			},
		)
		fmt.Println("closing grpc stream:")
		return nil

	case <-time.After(3 * time.Minute):
		fmt.Println("Game TimeOut")
	}
	return nil
}

func sendQuestionStream(stream pb.Game_StartGameServer, nextQuestionChannel chan int, closeChan chan error, result *pb.ResultStatus, timePerQuestion int32) {
	for i := 0; i < 10; i += 1 {
		q, ans, err := getRandomQuestion(i + 1)
		result.Ans = append(result.Ans, int32(*ans))
		result.Question = append(result.Question, q)
		result.ScoreOfEachQuestion = append(result.ScoreOfEachQuestion, 0)
		result.MarkedAns = append(result.MarkedAns, -1)
		if err != nil {
			closeChan <- err
			return
		}
		stream.Send(
			&pb.GameStatus{
				GameId: 1,
				Status: &pb.GameStatus_Question{
					Question: q,
				},
			},
		)
		select {
		case res := <-nextQuestionChannel:
			fmt.Println(res)
			result.MarkedAns[i] = int32(res)
			if res == *ans {
				result.Score += 10
				result.ScoreOfEachQuestion[i] = 10
			}
		case <-time.After(time.Second * time.Duration(timePerQuestion)):
			fmt.Println("Out of time :(")
		}
	}
	closeChan <- nil
}
func recieveAnswerStream(stream pb.Game_StartGameServer, nextQuestionChannel chan int) {
	for {
		userAction, err := stream.Recv()
		if err != nil {
			fmt.Println("Error in Starting Game", err)
			return
		}
		action := userAction.GetAction()
		switch action.(type) {
		case *pb.UserAction_GameRequest:
			fmt.Println("cant start new game while in already started game")
		case *pb.UserAction_QuestionAnswer:
			ans := userAction.GetQuestionAnswer().Answer
			fmt.Println("ans is ", ans)
			nextQuestionChannel <- int(ans)
		}
	}
}

func getRandomQuestion(QuestionNumber int) (*pb.Question, *int, error) {
	randAns := rand.Intn(4)
	logoModels, err := models.GetRandomLogo()
	if err != nil {
		return nil, nil, fmt.Errorf("Error In Random Question From SQL ")
	}
	switch rand.Intn(2) {
	case 0:
		q, err := getRandomImageQuestion(logoModels, randAns)
		if err != nil {
			return nil, nil, err
		}
		return &pb.Question{
			Id:             uint64(2),
			QuestionNumber: int32(QuestionNumber),
			Type: &pb.Question_ImageQuestion{
				ImageQuestion: q,
			},
		}, &randAns, nil

	case 1:
		q, err := getRandomNameQuestion(logoModels, randAns)
		if err != nil {
			return nil, nil, err
		}
		return &pb.Question{
			Id:             uint64(2),
			QuestionNumber: int32(QuestionNumber),
			Type: &pb.Question_NameQuestion{
				NameQuestion: q,
			},
		}, &randAns, nil
	default:
		return nil, nil, fmt.Errorf("Error in creating a random number")
	}
}

func getRandomImageQuestion(logoModels []models.Logo, randAns int) (*pb.ImageQuestion, error) {
	q := &pb.ImageQuestion{
		Id:   uint64(logoModels[randAns].ID),
		Name: &pb.LogoName{Name: logoModels[randAns].Name},
	}
	for i := 0; i < 4; i++ {
		logo, err := getPbLogoFromPath(logoModels[i].ImagePath)
		if err != nil {
			return nil, err
		}
		switch i {
		case 0:
			q.Image1 = logo
		case 1:
			q.Image2 = logo
		case 2:
			q.Image3 = logo
		case 3:
			q.Image4 = logo

		}
	}
	return q, nil
}

func getRandomNameQuestion(logoModels []models.Logo, randAns int) (*pb.NameQuestion, error) {

	q := &pb.NameQuestion{
		Id:    uint64(logoModels[randAns].ID),
		Name1: &pb.LogoName{Name: logoModels[0].Name},
		Name2: &pb.LogoName{Name: logoModels[1].Name},
		Name3: &pb.LogoName{Name: logoModels[2].Name},
		Name4: &pb.LogoName{Name: logoModels[3].Name},
	}
	logo, err := getPbLogoFromPath(logoModels[randAns].ImagePath)
	if err != nil {
		return nil, err
	}
	q.Image = logo
	return q, nil
}

func getPbLogoFromPath(path string) (*pb.Logo, error) {
	file, err := os.Open("../Logos/" + path)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Error While Reading Image in server")
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	// read file into bytes
	buffer := bufio.NewReader(file)
	buffer.Read(bytes)
	return &pb.Logo{Image: bytes}, nil

}
