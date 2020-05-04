package main

import (
	"context"
	"fmt"
	"foo/cmd/rpc/token/auth"
	"foo/cmd/rpc/token/greeter"
	"google.golang.org/grpc"
	"log"
)

func main() {
	a := auth.Auth{
		Account:  "douyacun",
		Password: "123456",
	}
	conn, err := grpc.Dial("localhost:4000", grpc.WithInsecure(), grpc.WithPerRPCCredentials(&a))
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
