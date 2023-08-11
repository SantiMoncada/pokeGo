package main

import (
	"net/http"
)

func setRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Home)
	mux.HandleFunc("/pokemons", Pokemons)
	mux.HandleFunc("/pokemon_details", PokemonDetails)

	dir := http.Dir("./static")
	fs := http.FileServer(dir)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	return mux
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
