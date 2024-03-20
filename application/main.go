package main

import (
	"antrein/bc-dashboard/application/grpc"
	"antrein/bc-dashboard/application/rest"
	"antrein/bc-dashboard/model/config"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	rest_app, err := rest.ApplicationDelegate(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Start gRPC server concurrently
	go func() {
		if err := grpc.StartServer(cfg); err != nil {
			log.Fatal(err)
		}
	}()

	if err = rest.StartServer(cfg, rest_app); err != nil {
		log.Fatal(err)
	}

}
