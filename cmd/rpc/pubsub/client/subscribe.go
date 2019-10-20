package main

import (
	"context"
	"fmt"
	"foo/cmd/rpc/pubsub/douyacun"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	cc, err := grpc.Dial(":12345", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial 错误：%v", err)
	}
	client := douyacun.NewPublisherClient(cc)
	stream, err := client.Subscribe(context.Background(), &douyacun.Topic{Name: "golang"})
	if err != nil {
		log.Fatalf("client.Subscribe 错误: %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Fatal(err)
		}
		fmt.Printf("新消息：%s\n", msg.String())
	}
}
