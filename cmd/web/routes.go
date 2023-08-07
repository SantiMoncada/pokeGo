package main

import (
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
		handleNotFound(w, r)
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAbout(w, r)
	default:
		handleNotFound(w, r)
	}
}
