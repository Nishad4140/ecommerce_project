package main

import (
	"log"

	"github.com/Nishad4140/ecommerce_project/pkg/config"
	"github.com/Nishad4140/ecommerce_project/pkg/di"
)

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI1(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
