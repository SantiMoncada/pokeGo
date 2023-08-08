package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseApi = "https://pokeapi.co/api/v2"

type Link struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Generations struct {
	Count    int64  `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Link `json:"results"`
}

func getPokeGenerations() (*Generations, error) {

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

func getPokeGeneration(i int64) (*Generation, error) {

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

	return &jsonResponse, nil
}

type Species struct {
	Name                 string `json:"name"`
	Id                   int64  `json:"id"`
	GenderRate           int64  `json:"gender_rate"`
	HatchCounter         int64  `json:"hatch_counter"`
	BaseHapiness         int64  `json:"base_happiness"`
	CaptureRate          int64  `json:"capture_rate"`
	HasGenderDifferences bool   `json:"has_gender_differences"`
	IsBaby               bool   `json:"is_baby"`
	IsLegendary          bool   `json:"is_legendary"`
	IsMythical           bool   `json:"is_mythical"`
	FormsSwitchable      bool   `json:"forms_switchable"`
	Color                Link   `json:"color"`
	EvolvesFromSpecies   Link   `json:"evolves_from_species"`
	GrowthRate           Link   `json:"growth_rate"`
	Generation           Link   `json:"generation"`
	Habitat              Link   `json:"habitat"`
	FlavorTextEntries    []struct {
		FlavorText string `json:"flavor_text"`
		Language   Link   `json:"language"`
		Version    Link   `json:"version"`
	} `json:"flavor_text_entries"`
}

func getPokemonSpecies(i int64) (*Species, error) {
	response, err := http.Get(fmt.Sprintf("%s/pokemon-species/%d/", baseApi, i))

	if err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var jsonResponse Species

	json.Unmarshal(responseData, &jsonResponse)

	return &jsonResponse, nil
}
