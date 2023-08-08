package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

func getHome(w http.ResponseWriter, _ *http.Request) {

	Generations, err := getPokeGenerations()

	if err != nil {
		fmt.Println(err)
	}

	parsedTemplate, err := template.ParseFiles("./templates/home.page.html", "./templates/base.layout.html")

	if err != nil {
		fmt.Println(err)
	}

	type templateDataType struct {
		Name string
		Id   string
	}

	var templateData []templateDataType

	for _, val := range Generations.Results {
		splited := strings.Split(val.Url, "/")
		id := splited[6]
		templateData = append(templateData, templateDataType{val.Name, id})
	}

	parsedTemplate.Execute(w, templateData)

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
