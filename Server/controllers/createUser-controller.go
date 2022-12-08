package controllers

import (
	"context"
	"fmt"
	"log"

	"github.com/J-Nokwal/Guess_The_Logo/Server/models"
	"github.com/J-Nokwal/Guess_The_Logo/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*emptypb.Empty, error) {

	log.Printf("Received: %v", in.GetUserName())
	user := models.User{UserName: in.GetUserName()}
	user.CreateUser()
	fmt.Println(server, user)
	header := metadata.Pairs("uid", fmt.Sprint(user.ID))
	grpc.SendHeader(ctx, header)
	return &emptypb.Empty{}, nil
}
