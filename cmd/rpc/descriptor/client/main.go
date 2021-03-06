package main

import (
	"context"
	"fmt"
	"foo/cmd/rpc/descriptor/greeter"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := greeter.NewGreeterClient(conn)

	reply, err := client.SayHello(context.Background(), &greeter.HelloRequest{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply.GetMessage())
}
