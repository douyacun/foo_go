package main

import (
	"context"
	"fmt"
	"foo/cmd/rpc/cs/greeter"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal("grp dial error:", err)
	}
	defer conn.Close()

	client := greeter.NewGreeterClient(conn)
	reply, err := client.SayHello(context.Background(), &greeter.HelloRequest{Name: "刘宁"})
	if err != nil {
		log.Fatal("sayHello error:", err)
	}
	fmt.Print(reply.Message)
}
