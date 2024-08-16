package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"hello/protobuf"
	"log"
)

/*
@Time : 2024/8/16 17:14
@Author : echo
@File : grpcClient
@Software: GoLand
@Description:
*/
func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := protobuf.NewGreeterClient(conn)
	res, err := c.SayHello(context.Background(), &protobuf.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	// 访问响应消息
	log.Printf("Greeting: %s", res.Message)

}
