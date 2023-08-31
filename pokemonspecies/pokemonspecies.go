package pokemonspecies

import (
	"encoding/json"

	"github.com/jeremyhager/pokeapi/pokeclient"
)

func Get(id string) (PokemonSpecies, error) {
	client := pokeclient.Init("pokemon-species")

	pokeByte, err := client.Get(id)
	if err != nil {
		return PokemonSpecies{}, err
	}
	var pokemon PokemonSpecies

	err = json.Unmarshal(pokeByte, &pokemon)
	if err != nil {
		return PokemonSpecies{}, err
	}
	return pokemon, nil
}
