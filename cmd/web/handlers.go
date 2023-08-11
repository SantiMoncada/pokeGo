package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

var templateCache = make(map[string]*template.Template)

func renderTemplate(templateName string) *template.Template {
	val, ok := templateCache[templateName]
	if ok {
		return val
	}

	parsedTemplate, err := template.ParseFiles(fmt.Sprintf("./templates/%s.page.html", templateName), "./templates/base.layout.html")

	if err != nil {
		log.Fatal(err)
	}

	templateCache[templateName] = parsedTemplate

	return parsedTemplate
}

func getHome(w http.ResponseWriter, _ *http.Request) {

	Generations, err := getPokeGenerations()

	if err != nil {
		fmt.Println(err)
	}

	parsedTemplate := renderTemplate("Home")

	parsedTemplate.Execute(w, Generations)
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

	parsedTemplate := renderTemplate("pokemons")

	parsedTemplate.Execute(w, templateData)
}

func getPokemonDetails(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	id := queryParams["id"]

	id2, _ := strconv.ParseInt(id[0], 10, 0)

	res, _ := getPokemonSpecies(id2)

	parsedTemplate := renderTemplate("pokemonDetails")

	parsedTemplate.Execute(w, res)
}

func handleNotFound(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "404 not found.", http.StatusNotFound)

}
