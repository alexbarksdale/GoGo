package main

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"grpc2/protos"
	"log"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"
)

var (
	client protos.BroadcastClient
	wg     *sync.WaitGroup
)

func init() {
	wg = &sync.WaitGroup{}
}

func connect(user *protos.User) error {
	var streamErr error

	stream, err := client.CreateStream(context.Background(), &protos.Connect{
		User:   user,
		Active: true,
	})

	if err != nil {
		return fmt.Errorf("connection failed: %v", err)
	}

	wg.Add(1)
	go func(str protos.Broadcast_CreateStreamClient) {
		defer wg.Done()

		for {
			msg, err := str.Recv()
			if err != nil {
				streamErr = fmt.Errorf("Error reading message: %v", err)
				break
			}

			fmt.Printf("%v : %s\n", msg.Id, msg.Content)
		}
	}(stream)

	return streamErr
}

func main() {
	timestamp := time.Now()
	done := make(chan int)

	name := flag.String("N", "User", "The name of the user")
	flag.Parse()

	id := sha256.Sum256([]byte(timestamp.String() + *name))

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}

	client = protos.NewBroadcastClient(conn)
	user := &protos.User{
		Id:   hex.EncodeToString(id[:]),
		Name: *name,
	}

	connect(user)

	wg.Add(1)

	go func() {
		defer wg.Done()

		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			msg := &protos.Message{
				Id:        user.Id,
				Content:   scanner.Text(),
				Timestamp: timestamp.String(),
			}

			_, err := client.BroadcastMessage(context.Background(), msg)
			if err != nil {
				fmt.Printf("Error sending msg: %v", err)
				break
			}
		}

	}()

	go func() {
		wg.Wait()
		close(done)
	}()

	<-done
}
