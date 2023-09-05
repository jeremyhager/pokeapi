package pokemonspecies

import (
	"encoding/json"
	"fmt"

	"github.com/jeremyhager/pokeapi/evolutionchains"
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

// Returns a flat array of all pokemon evolution species related to the current pokemon
func (species *PokemonSpecies) FlattenEvolutions(chain *evolutionchains.ChainLink) ([]PokemonSpecies, error) {
	evolutionChainPokemon := []PokemonSpecies{}

	for _, evolution := range chain.EvolvesTo {
		if len(evolution.EvolvesTo) < 0 {
			return evolutionChainPokemon, nil
		}

		species, err := Get(evolution.Species.Name)
		if err != nil {
			return evolutionChainPokemon, err
		}

		evolutionChainPokemon = append(evolutionChainPokemon, species)

		flattened, err := species.FlattenEvolutions(&evolution)
		if err != nil {
			return evolutionChainPokemon, err
		}
		evolutionChainPokemon = append(evolutionChainPokemon, flattened...)

	}
	return evolutionChainPokemon, nil

}

// Returns the base species evolution from an evolution chain
func (species *PokemonSpecies) GetBaseSpecies(chain *evolutionchains.ChainLink) (PokemonSpecies, error) {
	baseSpecies, err := Get(chain.Species.Name)
	fmt.Printf("base species Name: %v\n", baseSpecies.Name)
	if err != nil {
		return PokemonSpecies{}, err
	}
	return baseSpecies, nil
}
