package main

import (
	"flag"
	"github.com/ak4bento/clone-yourself/cmd"
	"github.com/ak4bento/clone-yourself/internal/server"
	"github.com/ak4bento/clone-yourself/internal/core"
)

func main() {
	core.InitDB() // Init DB on startup

	apiMode := flag.Bool("api", false, "Run in API mode")
	flag.Parse()

	if *apiMode {
		server.StartAPI()
	} else {
		cmd.Execute()
	}
}
