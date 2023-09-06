package pokeapi

type PokemonSpecies struct {
	ID                   int                      `json:"id"`
	Name                 string                   `json:"name"`
	Order                int                      `json:"order"`
	GenderRate           int                      `json:"gender_rate"`
	BaseHappiness        int                      `josn:"base_happiness"`
	IsBaby               bool                     `json:"is_baby"`
	IsLegendary          bool                     `json:"is_legendary"`
	IsMythical           bool                     `json:"is_mythical"`
	HatchCounter         int                      `json:"hatch_counter"`
	HasGenderDifferences bool                     `json:"has_gender_differences"`
	FormsSwitchable      bool                     `json:"forms_switchable"`
	GorthwRate           NamedAPIResource         `json:"growth_rate"`
	PokedexNumbers       []PokemonSpeciesDexEntry `json:"pokedex_numbers"`
	EggGroups            []NamedAPIResource       `json:"egg_groups"`
	Color                NamedAPIResource         `json:"color"`
	Shape                NamedAPIResource         `json:"shape"`
	EvolvesFromSpecies   NamedAPIResource         `json:"evolves_from_species"`
	EvolutionChain       APIResource              `json:"evolution_chain"`
	Habitat              NamedAPIResource         `json:"habitat"`
	Generation           NamedAPIResource         `json:"generation"`
	Names                []Name                   `json:"names"`
	PalParkEncounters    []PalParkEncounterArea   `json:"pal_park_encounters"`
	FlavorTextEntries    []FlavorText             `json:"flavor_text_entries"`
	FormDescriptions     []Description            `json:"form_descriptions"`
	Genera               []Genus                  `json:"genera"`
	Varieties            []PokemonSpeciesVariety  `json:"varieties"`
}

type PokemonSpeciesDexEntry struct {
	EntryNumber int              `json:"entry_number"`
	PokeDex     NamedAPIResource `json:"pokedex"`
}

type PalParkEncounterArea struct {
	BaseScore int              `json:"base_score"`
	Rate      int              `json:"rate"`
	Area      NamedAPIResource `json:"area"`
}

type Genus struct {
	Genus    string           `json:"genus"`
	Language NamedAPIResource `json:"language"`
}

type PokemonSpeciesVariety struct {
	IsDefault bool             `json:"is_default"`
	Pokemon   NamedAPIResource `json:"pokemon"`
}

// Returns a flat array of all pokemon evolution species related to the current pokemon
func (s *PokemonSpecies) FlattenEvolutions(chain *ChainLink) ([]PokemonSpecies, error) {
	evolutionChainPokemon := []PokemonSpecies{}

	for _, evolution := range chain.EvolvesTo {
		if len(evolution.EvolvesTo) < 0 {
			return evolutionChainPokemon, nil
		}

		species, err := GetSpecies(evolution.Species.Name)
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
func (s *PokemonSpecies) GetBaseSpecies(chain *ChainLink) (PokemonSpecies, error) {
	baseSpecies, err := GetSpecies(chain.Species.Name)
	// fmt.Printf("base species Name: %v\n", baseSpecies.Name)
	if err != nil {
		return PokemonSpecies{}, err
	}
	return baseSpecies, nil
}
