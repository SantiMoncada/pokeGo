package main

import (
	"fmt"
	"net/http"
)

func getHome(w http.ResponseWriter, _ *http.Request) {

	Generations, err := getPokemonSpecies(10)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, fmt.Sprintf("%#v\n", Generations))
	// parsedTemplate, err := template.ParseFiles("./templates/home.page.html", "./templates/base.layout.html")

	if err != nil {
		fmt.Println(err)
	}

	// err = parsedTemplate.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }

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
