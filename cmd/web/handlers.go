package main

import (
	"fmt"
	"net/http"
	"strconv"
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

func getPokemons(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()

	id := queryParams["id"]

	id2, _ := strconv.ParseInt(id[0], 10, 0)

	res, _ := getPokeGeneration(id2)

	type Specie struct {
		Name string
		Id   string
	}

	type templateDataType struct {
		Name    string
		Species []Specie
	}

	var templateData templateDataType

	templateData.Name = res.Name

	for _, val := range res.PokemonSpecies {
		splited := strings.Split(val.Url, "/")
		id := splited[6]
		templateData.Species = append(templateData.Species, Specie{val.Name, id})
	}

	parsedTemplate, err := template.ParseFiles("./templates/pokemons.page.html", "./templates/base.layout.html")

	if err != nil {
		fmt.Println(err)
	}

	parsedTemplate.Execute(w, templateData)
}

func getPokemonDetails(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	id := queryParams["id"]

	id2, _ := strconv.ParseInt(id[0], 10, 0)

	res, _ := getPokemonSpecies(id2)

	parsedTemplate, err := template.ParseFiles("./templates/pokemonDetails.page.html", "./templates/base.layout.html")

	if err != nil {
		fmt.Println(err)
	}

	parsedTemplate.Execute(w, res)
}

func handleNotFound(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "404 not found.", http.StatusNotFound)

}
