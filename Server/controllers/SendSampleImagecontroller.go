package controllers

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/J-Nokwal/Guess_The_Logo/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) SendSampleImage(ctx context.Context, in *emptypb.Empty) (*pb.Logo, error) {
	filePath := "../Logos/a-stars-alpinestars-vector-logo-400x400.png"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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
