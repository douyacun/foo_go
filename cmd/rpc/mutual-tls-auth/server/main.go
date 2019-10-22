package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"foo/cmd/rpc/cs/greeter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
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
	// 从磁盘加载证书
	cert, err := tls.LoadX509KeyPair("../cert/server.crt", "../cert/server.key")
	if err != nil {
		log.Fatalf("tls load key pair failed, %v",err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../cert/ca.crt")
	if err != nil {
		log.Fatalf("ca.crt read filed %v", err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append client certs")
	}
	// 初始化tcp通道
	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatal(err)
	}
	// 初始化TLS证书
	creds := credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{cert},
		ClientCAs:    certPool,
	})
	// 创建 gRPC 服务，并配置证书
	s := grpc.NewServer(grpc.Creds(creds))
	// 注册函数
	greeter.RegisterGreeterServer(s, new(server))
	// 启动服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
