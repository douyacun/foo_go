package main

import (
	"context"
	"foo/cmd/rpc2/greeter"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/rpc"
)

type server struct{}

func (p *server) SayHello(ctx context.Context, args *greeter.HelloRequest) (*greeter.HelloReply, error) {
	reply := &greeter.HelloReply{
		Message: "hello:" + args.Name,
	}
	return reply, nil
}

func main() {
	s := grpc.NewServer()
	greeter.RegisterGreeterServer(s, new(server))
	_ = rpc.RegisterName("arith", new(greeter.Arith))
	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	if err := s.Serve(listener); err != nil {
		log.Fatalf("serve listen error:%s", err)
	}
}
