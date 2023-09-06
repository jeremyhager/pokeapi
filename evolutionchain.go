package pokeapi

import (
	"reflect"
)

type EvolutionChain struct {
	ID              int               `json:"id"`
	BabyTriggerItem *NamedAPIResource `json:"baby_trigger_item,omitempty"`
	Chain           ChainLink         `json:"chain"`
}

type ChainLink struct {
	IsBaby           bool              `json:"is_baby,omitempty"`
	Species          NamedAPIResource  `json:"species"`
	EvolutionDetails []EvolutionDetail `json:"evolution_details,omitempty"`
	EvolvesTo        []ChainLink       `json:"evolves_to,omitempty"`
}

type EvolutionDetail struct {
	Item                  *NamedAPIResource `json:"item,omitempty"`
	Trigger               *NamedAPIResource `json:"trigger,omitempty"`
	Gender                int               `json:"gender,omitempty"`
	HeldItem              *NamedAPIResource `json:"held_item,omitempty"`
	KnownMove             *NamedAPIResource `json:"known_move,omitempty"`
	KnownMoveType         *NamedAPIResource `json:"known_move_type,omitempty"`
	Location              *NamedAPIResource `json:"location,omitempty"`
	MinLevel              int               `json:"min_level,omitempty"`
	MinHappiness          int               `json:"min_happiness,omitempty"`
	MinBeauty             int               `json:"min_beauty,omitempty"`
	MinAffection          int               `json:"min_affection,omitempty"`
	PartySpecies          *NamedAPIResource `json:"party_species,omitempty"`
	PartyType             *NamedAPIResource `json:"party_type,omitempty"`
	RelativePhysicalStats int               `json:"relative_physical_stats,omitempty"`
	TimeOfDay             string            `json:"time_of_day,omitempty"`
	TradeSpecies          *NamedAPIResource `json:"trade_species,omitempty"`
	TurnUpsideDown        bool              `json:"turn_upside_down,omitempty"`
}

// Returns a flat array of the evolution chain for the pokemon
func (c *ChainLink) Flatten() ([]ChainLink, error) {
	evolutionList := []ChainLink{}

	for _, evolution := range c.EvolvesTo {
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

func (e *EvolutionDetail) GetDetails() map[any]any {
	keyValue := make(map[any]any)

	v := reflect.ValueOf(*e)
	typeOfS := v.Type()
	// named := NamedAPIResource{}

	for i := 0; i < v.NumField(); i++ {
		fieldName := typeOfS.Field(i).Name
		value := v.Field(i).Interface()
		fieldType := v.Field(i).Type()
		if (value != 0) && (value != nil) && (value != "") && (fieldType != reflect.TypeOf(&NamedAPIResource{})) {
			// fmt.Printf("fieldType: %+v\t", fieldType)
			// fmt.Printf("fieldName: %s\tValue: %+v\n", fieldName, value)

			keyValue[fieldName] = value
		} else if (value != 0) && (value != nil) && (value != "") && (fieldType == reflect.TypeOf(&NamedAPIResource{})) {
			// fmt.Printf("fieldType: %+v\t", fieldType)
			// fmt.Printf("fieldName: %s\tValue: %+v\n", fieldName, value)
		}
	}
	return keyValue
}
