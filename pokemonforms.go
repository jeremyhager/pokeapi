package pokeapi

type PokemonForm struct {
	ID           string             `json:"id"`
	Name         string             `json:"name"`
	Order        int                `json:"order"`
	FormOrder    int                `json:"form_order"`
	IsDefault    bool               `json:"is_default"`
	IsBattleOnly bool               `json:"is_battle_only"`
	IsMega       bool               `json:"is_mega"`
	FormName     string             `json:"form_name"`
	Pokemon      NamedAPIResource   `json:"pokemon"`
	Types        []PokemonFormType  `json:"types"`
	Sprites      PokemonFormSprites `json:"sprites"`
	VersionGroup NamedAPIResource   `json:"version_group"`
	Names        []Name             `json:"names"`
	FormNames    []Name             `json:"form_names"`
}

type PokemonFormSprites struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
}

func GetPokemonForm(id string) (PokemonForm, error) {
	var form PokemonForm
	pokeapiUrl := client.urlBuilder(endpoints["pokemon-form"], id)
	err := client.Get(pokeapiUrl, &form)
	return form, err
}
