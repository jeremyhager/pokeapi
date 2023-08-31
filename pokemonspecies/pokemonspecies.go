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
	var species PokemonSpecies

	err = json.Unmarshal(pokeByte, &species)
	if err != nil {
		return PokemonSpecies{}, err
	}
	return species, nil
}
