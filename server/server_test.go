package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	"github.com/harshpreet93/lile_test"
	"github.com/lileio/lile/v2"
)

var s = LileTestServer{}
var cli lile_test.LileTestClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		lile_test.RegisterLileTestServer(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = lile_test.NewLileTestClient(lile.TestConn(addr))

	os.Exit(m.Run())
}
