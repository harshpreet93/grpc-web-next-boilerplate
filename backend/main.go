package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	"github.com/harshpreet93/grpc-web-next-boilerplate/backend/grpc_gen"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

// HelloServer asdf
type HelloServer struct {
}

// SayHello asdf
func (h *HelloServer) SayHello(ctx context.Context, req *grpc_gen.HelloRequest) (*grpc_gen.HelloReply, error) {
	log.Printf("received grpc request")
	return &grpc_gen.HelloReply{Message: fmt.Sprintf("hello %s", req.GetName())}, nil
}

func main() {
	var nFlag = flag.String("env", "dev", "help message for flag n")

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	greeterService := grpc_gen.NewGreeterService(&HelloServer{})
	grpc_gen.RegisterGreeterService(grpcServer, greeterService)

	r := gin.Default()

	wrappedGrpc := grpcweb.WrapServer(grpcServer)
	box := packr.New("My Box", "../frontend/out")
	// Greeter is the name of the service in the proto file
	r.Any("/Greeter/:call", func(c *gin.Context) {
		log.Printf("trying to route request")
		if wrappedGrpc.IsGrpcWebRequest(c.Request) {
			wrappedGrpc.ServeHTTP(c.Writer, c.Request)
		}
		// Fall back to other servers.
		// http.DefaultServeMux.ServeHTTP(c.Writer, c.Request)
	})
	r.StaticFS("/home", box)
	err := r.Run()
	if err != nil {
		log.Fatal("fail")
	}
}
