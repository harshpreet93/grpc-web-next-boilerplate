package main

import (
	_ "net/http/pprof"

	"github.com/harshpreet93/lile_test"
	"github.com/harshpreet93/lile_test/lile_test/cmd"
	"github.com/harshpreet93/lile_test/server"
	"github.com/lileio/fromenv"
	"github.com/lileio/lile/v2"
	"github.com/lileio/logr"
	"github.com/lileio/pubsub/v2"
	"github.com/lileio/pubsub/v2/middleware/defaults"
	"google.golang.org/grpc"
)

func main() {
	logr.SetLevelFromEnv()
	s := &server.LileTestServer{}

	lile.Name("lile_test")
	lile.Server(func(g *grpc.Server) {
		lile_test.RegisterLileTestServer(g, s)
	})

	pubsub.SetClient(&pubsub.Client{
		ServiceName: lile.GlobalService().Name,
		Provider:    fromenv.PubSubProvider(),
		Middleware:  defaults.Middleware,
	})

	cmd.Execute()
}
