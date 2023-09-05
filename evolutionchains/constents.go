package evolutionchains

import "github.com/jeremyhager/pokeapi/commonmodels"

type EvolutionChain struct {
	ID              int                            `json:"id"`
	BabyTriggerItem *commonmodels.NamedAPIResource `json:"baby_trigger_item,omitempty"`
	Chain           ChainLink                      `json:"chain"`
}

type ChainLink struct {
	IsBaby           bool                          `json:"is_baby,omitempty"`
	Species          commonmodels.NamedAPIResource `json:"species"`
	EvolutionDetails []EvolutionDetail             `json:"evolution_details,omitempty"`
	EvolvesTo        []ChainLink                   `json:"evolves_to,omitempty"`
}

type EvolutionDetail struct {
	Item                  *commonmodels.NamedAPIResource `json:"item,omitempty"`
	Trigger               *commonmodels.NamedAPIResource `json:"trigger,omitempty"`
	Gender                int                            `json:"gender,omitempty"`
	HeldItem              *commonmodels.NamedAPIResource `json:"held_item,omitempty"`
	KnownMove             *commonmodels.NamedAPIResource `json:"known_move,omitempty"`
	KnownMoveType         *commonmodels.NamedAPIResource `json:"known_move_type,omitempty"`
	Location              *commonmodels.NamedAPIResource `json:"location,omitempty"`
	MinLevel              int                            `json:"min_level,omitempty"`
	MinHappiness          int                            `json:"min_happiness,omitempty"`
	MinBeauty             int                            `json:"min_beauty,omitempty"`
	MinAffection          int                            `json:"min_affection,omitempty"`
	PartySpecies          *commonmodels.NamedAPIResource `json:"party_species,omitempty"`
	PartyType             *commonmodels.NamedAPIResource `json:"party_type,omitempty"`
	RelativePhysicalStats int                            `json:"relative_physical_stats,omitempty"`
	TimeOfDay             string                         `json:"time_of_day,omitempty"`
	TradeSpecies          *commonmodels.NamedAPIResource `json:"trade_species,omitempty"`
	TurnUpsideDown        bool                           `json:"turn_upside_down,omitempty"`
}
