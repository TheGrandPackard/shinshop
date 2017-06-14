package main

import (
	"log"
	"os"

	"github.com/Xackery/shinshop/webserver"
)

func main() {
	log.Println("Initializing...")
	err := webserver.Start("0.0.0.0:12345")
	if err != nil {
		log.Println("Error with webserver:", err.Error())
		os.Exit(1)
	}
}
