package main

import (
	"context"
	"fmt"
	"foo/cmd/rpc/pubsub/douyacun"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	cc, err := grpc.Dial(":12345", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := douyacun.NewPublisherClient(cc)
	for {
		resp, err := client.Publish(context.Background(), &douyacun.PublishRequest{
			Topic: &douyacun.Topic{
				Name: "golang",
			},
			Messages: &douyacun.PubsubMessage{
				Data: []byte("welcome!"),
			},
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(resp.MessageId)
		time.Sleep(time.Second)
	}
}
