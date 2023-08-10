package main

import (
	"net/http"
)

func setRoutes() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/pokemons", Pokemons)
	http.HandleFunc("/pokemon_details", PokemonDetails)
}

func Home(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getHome(w, r)
	default:
		handleNotFound(w, r)
	}
}

func Pokemons(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getPokemons(w, r)
	default:
		handleNotFound(w, r)
	}
}

func PokemonDetails(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getPokemonDetails(w, r)
	default:
		handleNotFound(w, r)
	}
}
