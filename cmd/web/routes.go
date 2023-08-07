package main

import (
	"fmt"
	"net/http"
)

func setRoutes() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
}

func Home(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getHome(w, r)
	case "POST":
		postHome(w, r)
	default:
		handleError(w, r)
	}
}

func getHome(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "hello world GET")

}

func postHome(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "hello world POST")

}

func About(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAbout(w, r)
	default:
		handleError(w, r)
	}
}

func getAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world about")

}

func handleError(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "404 not found.", http.StatusNotFound)
}
