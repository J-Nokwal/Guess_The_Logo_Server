package controllers

import (
	"log"
	"net"

	"github.com/J-Nokwal/Guess_The_Logo/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedGameServer
}

func Connect_grpc(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGameServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
