package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harshpreet93/grpc-web-next-boilerplate/backend/grpc_gen"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
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

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	greeterService := grpc_gen.NewGreeterService(&HelloServer{})
	grpc_gen.RegisterGreeterService(grpcServer, greeterService)

	r := gin.Default()
	wrappedGrpc := grpcweb.WrapServer(grpcServer)
	r.Any("/", func(c *gin.Context) {
		if wrappedGrpc.IsGrpcWebRequest(c.Request) {
			wrappedGrpc.ServeHTTP(c.Writer, c.Request)
		}
		// Fall back to other servers.
		http.DefaultServeMux.ServeHTTP(c.Writer, c.Request)
	})
}
