package main

import (
	"flag"
	"github.com/ak4bento/clone-yourself/cmd"
	"github.com/ak4bento/clone-yourself/internal/server"
)

func main() {
	apiMode := flag.Bool("api", false, "Run in API mode")
	flag.Parse()

	if *apiMode {
		server.StartAPI()
	} else {
		cmd.Execute()
	}
}
