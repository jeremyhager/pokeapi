package pokeapi

type Generation struct {
	ID             int                `json:"id"`
	Name           string             `json:"name"`
	Abilities      []NamedAPIResource `json:"abilities"`
	Names          []Name             `json:"names"`
	MainRegion     NamedAPIResource   `json:"main_region"`
	Moves          []NamedAPIResource `json:"moves"`
	PokemonSpecies []NamedAPIResource `json:"pokemon_species"`
	Types          []NamedAPIResource `json:"types"`
	VersionGroups  []NamedAPIResource `json:"version_groups"`
}

func GetGeneration(id string) (Generation, error) {
	var generation Generation
	pokeapiUrl := client.urlBuilder(endpoints["generation"], id)
	err := client.Get(pokeapiUrl, &generation)
	return generation, err
}
