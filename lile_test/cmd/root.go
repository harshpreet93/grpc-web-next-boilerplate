package cmd

import (
	"fmt"
	"os"

	"github.com/lileio/lile/v2"
)

var cfgFile string

var RootCmd = lile.BaseCommand("lile_test", "A gRPC based service")

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
