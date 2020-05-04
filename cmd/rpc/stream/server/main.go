package main

import (
	"foo/cmd/rpc/stream/greeter"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)
const Port = ":12345"
type server struct{}

func (s *server) RouteChat(stream greeter.Greeter_RouteChatServer) error {
	for{
		request, err := stream.Recv()
		if err != nil {
			// 客户端 断开当前连接
			if err == io.EOF {
				return nil
			}
			return err
		}
		err = stream.Send(&greeter.HelloReply{Message: "hello" + request.Name})
		if err != nil {
			return err
		}
	}
}

func main() {
	s := grpc.NewServer()
	greeter.RegisterGreeterServer(s, new(server))
	list, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatal(err)
	}
	if err = s.Serve(list); err!= nil {
		log.Fatal(err)
	}
}
