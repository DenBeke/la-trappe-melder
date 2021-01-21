package main

import (
	log "github.com/sirupsen/logrus"

	latrappemelder "github.com/DenBeke/la-trappe-melder"
)

func main() {

	config := latrappemelder.BuildConfigFromEnv()

	m, err := latrappemelder.New(config)
	if err != nil {
		log.Fatalf("couldn't init La Trappe Melder: %v", err)
	}

	m.Serve()

}
