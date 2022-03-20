package main

import (
	"fmt"
	"log"

	content "github.com/ellemouton/aperture-demo"
	"github.com/lightningnetwork/lnd/signal"
)

func main() {
	interceptor, err := signal.Intercept()
	if err != nil {
		log.Fatal(err)
	}

	s, err := content.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	defer s.Stop()

	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}

	<-interceptor.ShutdownChannel()
	fmt.Println("Received shutdown signal")

	if err := s.Stop(); err != nil {
		log.Fatal(err)
	}
}
