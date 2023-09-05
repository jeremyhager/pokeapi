package evolutionchains

import (
	"encoding/json"
	"reflect"

	"github.com/jeremyhager/pokeapi/commonmodels"
	"github.com/jeremyhager/pokeapi/pokeclient"
)

func Get(id string) (EvolutionChain, error) {
	client := pokeclient.Init("evolution-chain")

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

// Returns a flat array of the evolution chain for the pokemon
func (chain *ChainLink) Flatten() ([]ChainLink, error) {
	evolutionList := []ChainLink{}

	for _, evolution := range chain.EvolvesTo {
		if len(evolution.EvolvesTo) < 0 {
			return evolutionList, nil
		}

		evolutionList = append(evolutionList, evolution)

		flattened, err := evolution.Flatten()
		if err != nil {
			return evolutionList, err
		}
		evolutionList = append(evolutionList, flattened...)

	}
	return evolutionList, nil

}

func (detail *EvolutionDetail) GetDetails() map[any]any {
	keyValue := make(map[any]any)

	v := reflect.ValueOf(*detail)
	typeOfS := v.Type()
	// named := commonmodels.NamedAPIResource{}

	for i := 0; i < v.NumField(); i++ {
		fieldName := typeOfS.Field(i).Name
		value := v.Field(i).Interface()
		fieldType := v.Field(i).Type()
		if (value != 0) && (value != nil) && (value != "") && (fieldType != reflect.TypeOf(&commonmodels.NamedAPIResource{})) {
			// fmt.Printf("fieldType: %+v\t", fieldType)
			// fmt.Printf("fieldName: %s\tValue: %+v\n", fieldName, value)

			keyValue[fieldName] = value
		} else if (value != 0) && (value != nil) && (value != "") && (fieldType == reflect.TypeOf(&commonmodels.NamedAPIResource{})) {
			// fmt.Printf("fieldType: %+v\t", fieldType)
			// fmt.Printf("fieldName: %s\tValue: %+v\n", fieldName, value)
		}
	}
	return keyValue
}
