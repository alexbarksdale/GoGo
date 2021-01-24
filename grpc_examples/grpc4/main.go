package main

import (
	"log"
	"net"

	"github.com/alexbarksdale/GoGo/grpc_examples/grpc4/chat"
	"github.com/alexbarksdale/GoGo/grpc_examples/grpc4/pb"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	cs := chat.ChatServer{}

	gs := grpc.NewServer()

	pb.RegisterChatServiceServer(gs, &cs)

	if err := gs.Serve(listener); err != nil {
		log.Fatal("Failed to server gRPC server")
	}
}
