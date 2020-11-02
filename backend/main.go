package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/harshpreet93/grpc-web-next-boilerplate/backend/grpc_gen"
	"google.golang.org/grpc"
)

// HelloServer asdf
type HelloServer struct {
}

// SayHello asdf
func (h *HelloServer) SayHello(ctx context.Context, req *grpc_gen.HelloRequest) (*grpc_gen.HelloReply, error) {
	return &grpc_gen.HelloReply{Message: fmt.Sprintf("hello %s", req.GetName())}, nil
}

func main() {

	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	greeterService := grpc_gen.NewGreeterService(&HelloServer{})
	grpc_gen.RegisterGreeterService(grpcServer, greeterService)
	grpcServer.Serve(lis)
}
