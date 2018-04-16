package main

import (
	"log"
	"os"

	"github.com/derryx/go-microservice/internal/pkg/webserver"
)

var appName = "go-micro"

func setupLogger() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	setupLogger()
	log.Printf("Starting %v\n", appName)
	webserver.StartWebServer("8080") // NEW
}
