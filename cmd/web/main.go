package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	portNumber := os.Getenv("GO_SERVER_PORT")
	if portNumber == "" {
		portNumber = "8080"
	}

	host := os.Getenv("GO_SERVER_HOST")
	if os.Getenv("DOCKER_RUNNING") == "true" {
		host = "0.0.0.0"
	}
	if host == "" {
		host = "localhost"
	}

	mux := setRoutes()

	fmt.Println("Staring application on", host, ":", portNumber, portNumber)

	_ = http.ListenAndServe(host+":"+portNumber, mux)

}
