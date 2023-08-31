package pokemonspecies

import "github.com/jeremyhager/pokeapi/commonmodels"

type PokemonSpecies struct {
	ID                   int                             `json:"id"`
	Name                 string                          `json:"name"`
	Order                int                             `json:"order"`
	GenderRate           int                             `json:"gender_rate"`
	BaseHappiness        int                             `josn:"base_happiness"`
	IsBaby               bool                            `json:"is_baby"`
	IsLegendary          bool                            `json:"is_legendary"`
	IsMythical           bool                            `json:"is_mythical"`
	HatchCounter         int                             `json:"hatch_counter"`
	HasGenderDifferences bool                            `json:"has_gender_differences"`
	FormsSwitchable      bool                            `json:"forms_switchable"`
	GorthwRate           commonmodels.NamedAPIResource   `json:"growth_rate"`
	PokedexNumbers       []PokemonSpeciesDexEntry        `json:"pokedex_numbers"`
	EggGroups            []commonmodels.NamedAPIResource `json:"egg_groups"`
	Color                commonmodels.NamedAPIResource   `json:"color"`
	Shape                commonmodels.NamedAPIResource   `json:"shape"`
	EvolvesFromSpecies   commonmodels.NamedAPIResource   `json:"evolves_from_species"`
	EvolutionChain       commonmodels.APIResource        `json:"evolution_chain"`
	Habitat              commonmodels.NamedAPIResource   `json:"habitat"`
	Generation           commonmodels.NamedAPIResource   `json:"generation"`
	Names                []commonmodels.Name             `json:"names"`
	PalParkEncounters    []PalParkEncounterArea          `json:"pal_park_encounters"`
	FlavorTextEntries    []commonmodels.FlavorText       `json:"flavor_text_entries"`
	FormDescriptions     []commonmodels.Description      `json:"form_descriptions"`
	Genera               []Genus                         `json:"genera"`
	Varieties            []PokemonSpeciesVariety         `json:"varieties"`
}

type PokemonSpeciesDexEntry struct {
	EntryNumber int                           `json:"entry_number"`
	PokeDex     commonmodels.NamedAPIResource `json:"pokedex"`
}

type PalParkEncounterArea struct {
	BaseScore int                           `json:"base_score"`
	Rate      int                           `json:"rate"`
	Area      commonmodels.NamedAPIResource `json:"area"`
}

type Genus struct {
	Genus    string                        `json:"genus"`
	Language commonmodels.NamedAPIResource `json:"language"`
}

type PokemonSpeciesVariety struct {
	IsDefault bool                          `json:"is_default"`
	Pokemon   commonmodels.NamedAPIResource `json:"pokemon"`
}
