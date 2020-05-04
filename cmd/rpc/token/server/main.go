package main

import (
	"context"
	"fmt"
	"foo/cmd/rpc/token/auth"
	"foo/cmd/rpc/token/greeter"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
)

const Port = ":4000"

type server struct{ auth *auth.Auth }

func (p *server) SayHello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloReply, error) {
	reply := &greeter.HelloReply{
		Message: "hello " + req.GetName(),
	}
	return reply, nil
}

func main() {
	s := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_auth.UnaryServerInterceptor(check),
		grpc_recovery.UnaryServerInterceptor(),
	)))
	greeter.RegisterGreeterServer(s, &server{})
	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}

func check(ctx context.Context) (context.Context, error){
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing credentials")
	}
	var (
		account, password string
	)
	if val, ok := md["account"]; ok {
		account = val[0]
	}
	if val, ok := md["password"]; !ok {
		password = val[0]
	}
	fmt.Println(account, password)
	return nil, nil
}