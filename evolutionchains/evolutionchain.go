package evolutionchains

import (
	"encoding/json"

	"github.com/jeremyhager/pokeapi/pokeclient"
)

func Get(id string) (EvolutionChain, error) {
	client := pokeclient.Init("evolution-chain")

	chainByte, err := client.Get(id)
	if err != nil {
		return EvolutionChain{}, err
	}
	var chain EvolutionChain

	err = json.Unmarshal(chainByte, &chain)
	if err != nil {
		return EvolutionChain{}, err
	}
	return chain, nil
}
