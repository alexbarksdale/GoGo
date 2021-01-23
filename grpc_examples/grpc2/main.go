package main

import (
	"context"
	"grpc2/protos"
	"log"
	"net"
	"os"
	"sync"

	"google.golang.org/grpc"
	glog "google.golang.org/grpc/grpclog"
)

var grpcLog glog.LoggerV2

func init() {
	grpcLog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
}

type Connection struct {
	stream protos.Broadcast_CreateStreamServer
	id     string
	active bool
	err    chan error
}

type Server struct {
	Connection []*Connection
}

func (s *Server) CreateStream(pconn *protos.Connect, stream protos.Broadcast_CreateStreamServer) error {
	conn := &Connection{
		stream: stream,
		id:     pconn.User.Id,
		active: true,
		err:    make(chan error),
	}

	s.Connection = append(s.Connection, conn)

	return <-conn.err
}

func (s *Server) BroadcastMessage(ctx context.Context, msg *protos.Message) (*protos.Close, error) {
	wg := sync.WaitGroup{}

	done := make(chan int)

	for _, conn := range s.Connection {
		wg.Add(1)

		go func(msg *protos.Message, conn *Connection) {
			defer wg.Done()

			if conn.active {
				err := conn.stream.Send(msg)
				grpcLog.Info("Sending message to: ", conn.stream)

				if err != nil {
					grpcLog.Errorf("Error with stream: %s - Error: %v", conn.stream, err)
					conn.active = false
					conn.err <- err
				}
			}
		}(msg, conn)
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	<-done
	return &protos.Close{}, nil
}

func main() {
	var connections []*Connection

	server := &Server{connections}

	gs := grpc.NewServer()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	grpcLog.Info("Starting server on port :8080")

	protos.RegisterBroadcastServer(gs, server)
	gs.Serve(listener)
}
