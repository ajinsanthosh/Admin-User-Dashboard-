package main

import (
	config "example/pkg/config"
	"example/pkg/di"
	"log"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	server, err := di.InitializeAPI(config)
	if err != nil {
		log.Fatal("can't start server ", err)
	} else {
		server.Start()
	}

}
