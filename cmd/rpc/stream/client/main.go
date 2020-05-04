package main

import (
	"context"
	"fmt"
	"foo/cmd/rpc/stream/greeter"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

const Port = ":12345"

func main() {
	cc, err := grpc.Dial(Port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := greeter.NewGreeterClient(cc)
	stream, err := client.RouteChat(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			if err := stream.Send(&greeter.HelloRequest{Name: "world"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Fatal(err)
		}
		fmt.Println(reply.Message)
	}
}
