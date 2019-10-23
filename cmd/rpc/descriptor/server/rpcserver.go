package main

import (
	"context"
	"foo/cmd/rpc/cs/greeter"
	"google.golang.org/grpc"
	"log"
	"net"
)

const Port = ":4000"

type server struct{}

func (p *server) SayHello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloReply, error) {
	reply := &greeter.HelloReply{
		Message: "hello " + req.GetName(),
	}
	return reply, nil
}

func main() {
	s := grpc.NewServer()
	greeter.RegisterGreeterServer(s, new(server))
	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
