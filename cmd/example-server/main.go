package main

import (
	"flag"
	"fmt"
	"http-server-example/config"
	"http-server-example/server"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	defaultServiceName = "service"
)

func main() {
	name := flag.String("name", defaultServiceName, "the proper name of the service for reporting")
	cfg, err := config.Read(*name)
	if err != nil {
		log.Panic("Unable to parse configuration from env vars error: ", err)
	}

	app, err := server.New(cfg)
	if err != nil {
		log.Panic("Unable to setup the clients properly error: ", err)
	}

	router := server.ConfigureRouter(app)

	log.Println("Server up and listening on port: ", cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router); err != nil {
		log.Panicf("Unable to setup server on port: %d", cfg.Port)
	}
}
