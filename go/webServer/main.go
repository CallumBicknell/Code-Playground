package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/CallumBicknell/go-webServer/config"
	"github.com/CallumBicknell/go-webServer/db"
	"github.com/CallumBicknell/go-webServer/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()

	config.Init(*environment)

	db.Init()
	server.Init()
}
