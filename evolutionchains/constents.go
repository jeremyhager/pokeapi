package evolutionchains

import "github.com/jeremyhager/pokeapi/commonmodels"

type EvolutionChain struct {
	ID              int                           `json:"id"`
	BabyTriggerItem commonmodels.NamedAPIResource `json:"baby_trigger_item"`
	Chain           ChainLink                     `json:"chain"`
}

type ChainLink struct {
	IsBaby           bool                          `json:"is_baby"`
	Species          commonmodels.NamedAPIResource `json:"species"`
	EvolutionDetails []EvolutionDetail             `json:"evolution_details"`
	EvolvesTo        []ChainLink                   `json:"evolves_to"`
}

type EvolutionDetail struct {
	Item                  commonmodels.NamedAPIResource `json:"item"`
	Trigger               commonmodels.NamedAPIResource `json:"trigger"`
	Gender                int                           `json:"gender"`
	HeldItem              commonmodels.NamedAPIResource `json:"held_item"`
	KnownMove             commonmodels.NamedAPIResource `json:"known_move"`
	KnownMoveType         commonmodels.NamedAPIResource `json:"known_move_type"`
	Location              commonmodels.NamedAPIResource `json:"location"`
	MinLevel              int                           `json:"min_level"`
	MinHappiness          int                           `json:"min_happiness"`
	MinBeauty             int                           `json:"min_beauty"`
	MinAffection          int                           `json:"min_affection"`
	PartySpecies          commonmodels.NamedAPIResource `json:"party_species"`
	PartyType             commonmodels.NamedAPIResource `json:"party_type"`
	RelativePhysicalStats int                           `json:"relative_physical_stats"`
	TimeOfDay             string                        `json:"time_of_day"`
	TradeSpecies          commonmodels.NamedAPIResource `json:"trade_species"`
	TurnUpsideDown        bool                          `json:"turn_upside_down"`
}
