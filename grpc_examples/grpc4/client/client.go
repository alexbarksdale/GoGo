package main

import (
	"context"
	"log"

	"github.com/alexbarksdale/GoGo/grpc_examples/grpc4/pb"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	cs := pb.NewChatServiceClient(conn)

	msg := pb.Message{
		Body: "Hello from client!",
	}

	resp, err := cs.SayHello(context.Background(), &msg)
	if err != nil {
		log.Fatalf("Failed to say hello: %s", err)
	}

	log.Printf("Response from Server: %s", resp.Body)
}
