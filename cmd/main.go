package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/realHoangHai/authenticator/config"
	"github.com/realHoangHai/authenticator/pkg/graceful"
	"github.com/realHoangHai/authenticator/pkg/log"
	"os"
)

// VERSION Usage: go build -ldflags "-X main.VERSION=x.x.x"
var VERSION = "1.0"

var (
	file string
)

func showHelp() {
	fmt.Printf("Usage:%s {params}\n", os.Args[0])
	fmt.Println("      -c {config file}")
	fmt.Println("      -h (show help info)")
}

func parse() bool {
	flag.StringVar(&file, "c", "config/dev.toml", "config file")
	help := flag.Bool("h", false, "help info")
	flag.Parse()

	if !config.Load(file) {
		return false
	}

	if *help {
		showHelp()
		return false
	}
	return true
}

// @title authenticator
// @version VERSION
// @description Go Authenticator.
// @termsOfService http://swagger.io/terms/
// @basePath /
// @license.name MIT
// @license.url http://github.com/realHoangHai/authenticator/blob/master/LICENSE
// @contact.name realHoangHai
// @contact.email aflyingpenguin2lth@gmail.com
func main() {
	if !parse() {
		showHelp()
		os.Exit(-1)
	}

	log.Init("debug")

	// init	server
	server, err := initServer(context.Background())
	if err != nil {
		log.F("initialize server: %v", err)
	}

	g := graceful.NewManager()
	g.AddRunningJob(func(ctx context.Context) error {
		return server.Run(ctx)
	})

	<-g.Done()
}
