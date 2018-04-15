package main

import (
	"fmt"
	"go-microservice/internal/pkg/webserver"
)

var appName = "go-micro"

func main() {
	fmt.Printf("Starting %v\n", appName)
	webserver.StartWebServer("8080") // NEW
}
