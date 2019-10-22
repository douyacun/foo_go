package main

import (
	"context"
	"foo/cmd/rpc/token/auth"
	"foo/cmd/rpc/token/greeter"
	"google.golang.org/grpc"
	"log"
	"net"
)

const Port = ":4000"

type server struct{ auth *auth.Auth }

func (p *server) SayHello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloReply, error) {
	if err := p.auth.Check(ctx); err != nil {
		return nil, err
	}
	reply := &greeter.HelloReply{
		Message: "hello " + req.GetName(),
	}
	return reply, nil
}

func main() {
	s := grpc.NewServer()
	greeter.RegisterGreeterServer(s, &server{})
	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
