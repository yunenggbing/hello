package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"

	"hello/protobuf"
	"net"
)

/*
@Time : 2024/8/16 15:06
@Author : echo
@File : grpcServer
@Software: GoLand
@Description: grpc服务端
*/
type server struct {
	protobuf.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *protobuf.HelloRequest) (*protobuf.HelloReply, error) {
	return &protobuf.HelloReply{Message: "Hello " + in.Name}, nil
}
func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protobuf.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
