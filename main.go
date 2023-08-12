package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	portNumber := os.Getenv("PORT")
	if portNumber == "" {
		portNumber = "8080"

	}

	host := os.Getenv("HOST")

	mux := setRoutes()

	fmt.Println("Staring application on", ":", portNumber)

	_ = http.ListenAndServe(host+":"+portNumber, mux)

}
