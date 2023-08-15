package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const baseApi = "https://pokeapi.co/api/v2"

const timeToExpire = 168 * time.Hour

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
	Count     int64  `json:"count"`
	Next      string `json:"next"`
	Previous  string `json:"previous"`
	Results   []Link `json:"results"`
	timestamp time.Time
}

var generationsCache *Generations = nil

func getPokeGenerations() (*Generations, error) {

	if generationsCache != nil {
		timeFetched := generationsCache.timestamp
		expirationDate := timeFetched.Add(timeToExpire)

		if time.Now().Before(expirationDate) {
			return generationsCache, nil
		}

		generationsCache = nil
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

	jsonResponse.timestamp = time.Now()

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
	timestamp time.Time
}

var generationDataCache = make(map[int64]*Generation)

func getPokeGeneration(i int64) (*Generation, error) {

	generationCache, ok := generationDataCache[i]

	if ok {
		timeFetched := generationCache.timestamp
		expirationDate := timeFetched.Add(timeToExpire)

		if time.Now().Before(expirationDate) {
			return generationCache, nil
		}

		delete(generationDataCache, i)
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

	jsonResponse.timestamp = time.Now()

	generationDataCache[i] = &jsonResponse

	return &jsonResponse, nil
}

type Species struct {
	Name           string `json:"name"`
	Id             int64  `json:"id"`
	Weight         int64  `json:"weight"`
	Height         int64  `json:"height"`
	BaseExperience int64  `json:"base_experience"`
	Sprites        struct {
		FrontDefault string `json:"front_default"`
		BackDefault  string `json:"back_default"`
		FrontShiny   string `json:"front_shiny"`
		BackShiny    string `json:"back_shiny"`
	} `json:"sprites"`
	GameIndices []struct {
		GameIndex int64 `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		}
	} `json:"game_indices"`
	timestamp time.Time
}

var pokemonsCache = make(map[int64]*Species)

func getPokemonSpecies(i int64) (*Species, error) {

	pokemonCache, ok := pokemonsCache[i]

	if ok {
		timeFetched := pokemonCache.timestamp
		expirationDate := timeFetched.Add(timeToExpire)

		if time.Now().Before(expirationDate) {
			return pokemonCache, nil
		}

		delete(pokemonsCache, i)
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

	jsonResponse.timestamp = time.Now()

	pokemonsCache[i] = &jsonResponse

	return &jsonResponse, nil
}
