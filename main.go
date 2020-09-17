package main

import (
	"flag"
	"log"

	"github.com/Clever/analytics-latency-config-service/config"
	"github.com/Clever/analytics-latency-config-service/controller"
	"github.com/Clever/analytics-latency-config-service/gen-go/server"
	"github.com/Clever/wag/swagger"
)

func main() {
	addr := flag.String("addr", ":6723", "Address to listen at")
	flag.Parse()
	config.Init()

	swagger.InitCustomFormats()

	controller, err := controller.New()
	if err != nil {
		log.Fatal(err)
	}

	s := server.New(controller, *addr)
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}

	log.Println("analytics-latency-config-service exited without error")
}
