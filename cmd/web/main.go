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

	mux := setRoutes()

	fmt.Println("Staring application on port", portNumber)

	_ = http.ListenAndServe("localhost:"+portNumber, mux)

}
