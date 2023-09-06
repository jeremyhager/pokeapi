package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	scheme   = "https"
	host     = "pokeapi.co"
	rootPath = "/api/v2"
	retries  = 3
)

// type Pokeapi interface {
// 	Init(endpoint string) *Client
// 	GetEvolutions(id string) (EvolutionChain, error)
// 	GetSpecies(id string) (PokemonSpecies, error)
// 	GetPokemon(id string) (Pokemon, error)
// }

type Client struct {
	requestUrl string
}

// Init to create client
//
// Example:
//
//	client := pokeclint.Init("pokemon")
func Init(endpoint string) *Client {
	POKEURL := fmt.Sprintf("%v://%v/%v", scheme, host, rootPath)
	reqUrl := fmt.Sprintf("%v/%v", POKEURL, endpoint)
	client := &Client{
		requestUrl: reqUrl,
	}

	return client
}

func (c *Client) Get(id string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%v/%v", c.requestUrl, id))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}

func GetSpecies(id string) (PokemonSpecies, error) {
	client := Init("pokemon-species")

	pokeByte, err := client.Get(id)
	if err != nil {
		return PokemonSpecies{}, err
	}
	var species PokemonSpecies

	err = json.Unmarshal(pokeByte, &species)
	if err != nil {
		return PokemonSpecies{}, err
	}
	return species, nil
}

func GetPokemon(id string) (Pokemon, error) {
	client := Init("pokemon")

	pokeByte, err := client.Get(id)
	if err != nil {
		return Pokemon{}, err
	}
	var pokemon Pokemon

	err = json.Unmarshal(pokeByte, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}

func GetEvolutions(id string) (EvolutionChain, error) {
	client := Init("evolution-chain")

	chainByte, err := client.Get(id)
	if err != nil {
		return EvolutionChain{}, err
	}
	var chain EvolutionChain
	// err = json.Compact()

	err = json.Unmarshal(chainByte, &chain)
	if err != nil {
		return EvolutionChain{}, err
	}
	return chain, nil
}

func GetGenerations(id string) (Generations, error) {
	client := Init("generation")

	chainByte, err := client.Get(id)
	if err != nil {
		return Generations{}, err
	}
	var generation Generations
	// err = json.Compact()

	err = json.Unmarshal(chainByte, &generation)
	if err != nil {
		return Generations{}, err
	}
	return generation, nil
}
