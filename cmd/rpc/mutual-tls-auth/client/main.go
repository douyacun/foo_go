package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"foo/cmd/rpc/cs/greeter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

func main() {
	// 从磁盘加载证书
	certificate, err := tls.LoadX509KeyPair("../cert/client.crt", "../cert/client.key")
	if err != nil {
		log.Fatalf("tls load key pair failed, %v", err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../cert/ca.crt")
	if err != nil {
		log.Fatalf("ca read failed, %v", err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("cert append pool failed")
	}
	// 初始化TLS证书
	creds := credentials.NewTLS(&tls.Config{
		ServerName:   "douyacun.com", // NOTE: this is required!
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	})
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
