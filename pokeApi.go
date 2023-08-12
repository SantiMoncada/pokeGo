package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const baseApi = "https://pokeapi.co/api/v2"

type Link struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Id   string
}

func getIdFromLink(link *Link) {
	splited := strings.Split(link.Url, "/")
	id := splited[6]
	link.Id = id
}

type Generations struct {
	Count    int64  `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Link `json:"results"`
}

var generationsCache *Generations = nil

func getPokeGenerations() (*Generations, error) {

	if generationsCache != nil {
		return generationsCache, nil
	}

	response, err := http.Get(fmt.Sprintf("%s/generation", baseApi))

	if err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var jsonResponse Generations

	json.Unmarshal(responseData, &jsonResponse)

	for i := range jsonResponse.Results {
		getIdFromLink(&jsonResponse.Results[i])
	}

	generationsCache = &jsonResponse
	return &jsonResponse, nil
}

type Generation struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	MainRegion     Link   `json:"main_region"`
	PokemonSpecies []Link `json:"pokemon_species"`
	Types          []Link `json:"types"`
	Moves          []Link `json:"moves"`
	VersionGroups  []Link `json:"version_groups"`
	Names          []struct {
		Name     string `json:"name"`
		Language Link   `json:"language"`
	} `json:"names"`
}

var generationDataCache = make(map[int64]*Generation)

func getPokeGeneration(i int64) (*Generation, error) {

	val, ok := generationDataCache[i]

	if ok {
		return val, nil
	}

	response, err := http.Get(fmt.Sprintf("%s/generation/%d/", baseApi, i))

	if err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var jsonResponse Generation

	json.Unmarshal(responseData, &jsonResponse)

	for _, val := range jsonResponse.Names {
		getIdFromLink(&val.Language)
	}

	generationDataCache[i] = &jsonResponse

	return &jsonResponse, nil
}

type Species struct {
	Name           string `json:"name"`
	Id             int64  `json:"id"`
	Weight         int64  `json:"weight"`
	BaseExperience int64  `json:"base_experience"`
	Sprites        struct {
		FrontDefault string `json:"front_default"`
		BackDefault  string `json:"back_default"`
		FrontShiny   string `json:"front_shiny"`
		BackShiny    string `json:"back_shiny"`
	} `json:"sprites"`
}

var pokemonCache = make(map[int64]*Species)

func getPokemonSpecies(i int64) (*Species, error) {

	val, ok := pokemonCache[i]

	if ok {
		return val, nil
	}

	response, err := http.Get(fmt.Sprintf("%s/pokemon/%d/", baseApi, i))

	if err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var jsonResponse Species

	json.Unmarshal(responseData, &jsonResponse)

	pokemonCache[i] = &jsonResponse

	return &jsonResponse, nil
}
