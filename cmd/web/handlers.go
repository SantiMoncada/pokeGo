package main

import (
	"fmt"
	"net/http"
)

func getHome(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "hello world GET")

}

func postHome(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "hello world POST")

}

func getAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world about")

}

func handleNotFound(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "404 not found.", http.StatusNotFound)
}
