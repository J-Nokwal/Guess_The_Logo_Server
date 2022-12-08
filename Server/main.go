package main

import (
	"github.com/J-Nokwal/Guess_The_Logo/Server/controllers"
	_ "github.com/J-Nokwal/Guess_The_Logo/Server/models"
)

func main() {
	var port string = ":50051"
	controllers.Connect_grpc(port)

}
