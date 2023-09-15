package pokeapi

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

// Get evolution chain by ID
//
//	chain, _ := GetEvolutionChain("80")
func GetEvolutionChain(id string) (EvolutionChain, error) {
	var evolution EvolutionChain
	pokeapiUrl := client.urlBuilder(endpoints["evolution-chain"], id)
	err := client.Get(pokeapiUrl, &evolution)
	return evolution, err
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
