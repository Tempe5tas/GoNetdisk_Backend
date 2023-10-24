//go:build go1.16
// +build go1.16

package main

import (
	cmd2 "go-netdisk/cmd"
	"go-netdisk/internal/version"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-netdisk"
	app.Usage = "A simple net-disk service"
	app.Version = version.Version
	app.Commands = []*cli.Command{
		cmd2.Web,
		cmd2.Migrate,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}
}
