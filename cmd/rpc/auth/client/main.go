package main

import (
	"context"
	"fmt"
	"foo/cmd/rpc/cs/greeter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	// 初始化TLS证书
	creds, err := credentials.NewClientTLSFromFile("../cert/server.crt", "server.douyacun.com")
	if err != nil {
		log.Fatal(err)
	}
	// 拨号通信
	conn, err := grpc.Dial("localhost:4000", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := greeter.NewGreeterClient(conn)

	reply, err := client.SayHello(context.Background(), &greeter.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply.GetMessage())
}
