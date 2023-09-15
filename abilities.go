package pokeapi

type Ability struct {
	ID                int                   `json:"id"`
	Name              string                `json:"name"`
	IsMainSeries      bool                  `json:"is_main_series"`
	Generation        NamedAPIResource      `json:"generation"`
	Names             []Name                `json:"names"`
	EffectEntries     []VerboseEffect       `json:"effect_entries"`
	EffectChanges     []AbilityEffectChange `json:"effect_changes"`
	FlavorTextEntries []AbilityFlavorText   `json:"flavor_text_entries"`
	Pokemon           []AbilityPokemon      `json:"pokemon"`
}

type AbilityEffectChange struct {
	EffectEntries []VerboseEffect  `json:"effect_entries"`
	VersionGroup  NamedAPIResource `json:"version_group"`
}

type AbilityFlavorText struct {
	FlavorText   string           `json:"flavor_text"`
	Language     NamedAPIResource `json:"language"`
	VersionGroup NamedAPIResource `json:"version_group"`
}

type AbilityPokemon struct {
	IsHidden bool             `json:"is_hidden"`
	Slot     int              `json:"slot"`
	Pokemon  NamedAPIResource `json:"pokemon"`
}

func GetAbility(id string) (Ability, error) {
	var ability Ability
	pokeapiUrl := client.urlBuilder(endpoints["ability"], id)
	err := client.Get(pokeapiUrl, &ability)
	return ability, err
}
