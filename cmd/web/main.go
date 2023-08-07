package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	//get port number
	portNumber := os.Getenv("GO_SERVER_PORT")
	if portNumber == "" {
		portNumber = "8080"
	}

	setRoutes()

	fmt.Println("Staring application on port", portNumber)

	_ = http.ListenAndServe(":"+portNumber, nil)

}
