package main

import (
	"net"

	"github.com/alexbarksdale/GoGo/micro/currency/pb"
	"github.com/alexbarksdale/GoGo/micro/currency/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	cs := server.NewCurrencyServer(log)

	pb.RegisterCurrencyServer(gs, cs)

	// Register reflection service on gRPC server.
	// Do not use on production.
	reflection.Register(gs)

	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		panic(err)
	}

	if err := gs.Serve(listener); err != nil {
		panic(err)
	}
}
