package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/facebook/ent/entc"
	"github.com/facebook/ent/entc/gen"
	"github.com/facebook/ent/schema/field"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gobuffalo/packr/v2"
	"github.com/harshpreet93/grpc-web-next-boilerplate/backend/config"
	"github.com/harshpreet93/grpc-web-next-boilerplate/backend/ent"
	"github.com/harshpreet93/grpc-web-next-boilerplate/backend/ent/migrate"
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

func initDBConn(conf config.Config) *ent.Client {
	client, err := ent.Open("mysql",
		fmt.Sprintf("app:%s@tcp(localhost:3306)/app", conf.DBPass))
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	return client
}

func migrateDB(conn *ent.Client) error {
	ctx := context.Background()
	err := conn.Schema.Create(
		ctx,
		migrate.WithDropIndex(false),
		migrate.WithDropColumn(false),
	)
	return err
}

func generate() {
	err := entc.Generate("./ent/schema", &gen.Config{
		Header: "// Your Custom Header",
		IDType: &field.TypeInfo{Type: field.TypeInt},
	})
	if err != nil {
		log.Fatal("running ent codegen:", err)
	}
}

func main() {
	shouldGenerate := flag.Bool("generate", false, "Set generate to true to regenerate DB model code")

	var env = flag.String("env", "dev", "help message for flag n")
	flag.Parse()
	if *shouldGenerate {
		log.Printf("Generating DB code")
		generate()
		os.Exit(0)
	}
	conf, err := config.New(*env)

	if err != nil {
		log.Fatalf("Config could not be initialized %+v", err)
	}

	conn := initDBConn(conf)
	defer conn.Close()
	err = migrateDB(conn)

	if err != nil {
		log.Fatalf("Error migrating DB %+v", err)
	}

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
	err = r.Run()
	if err != nil {
		log.Fatal("fail")
	}
}
