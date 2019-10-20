package main

import (
	"context"
	"foo/cmd/rpc/pubsub/douyacun"
	"github.com/docker/docker/pkg/pubsub"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const Port = ":12345"

type server struct {
	Pub *pubsub.Publisher
}

func (s *server) Publish(c context.Context, pub *douyacun.PublishRequest) (*douyacun.PublishResponse, error) {
	s.Pub.Publish(pub)
	return &douyacun.PublishResponse{MessageId: 1}, nil
}

func (s *server) Subscribe(req *douyacun.Topic, stream douyacun.Publisher_SubscribeServer) error {
	ch := s.Pub.SubscribeTopic(func(v interface{}) bool {
		if req.GetName() == "" {
			return true
		}
		if it, ok := v.(*douyacun.PublishRequest); ok {
			if it.Topic.GetName() == req.GetName() {
				return true
			}
		}
		return false
	})
	for v := range ch {
		if pub, ok := v.(*douyacun.PublishRequest); ok {
			if err := stream.Send(pub.GetMessages()); err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	douyacun.RegisterPublisherServer(s, &server{Pub: pubsub.NewPublisher(100*time.Millisecond, 10)})
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
