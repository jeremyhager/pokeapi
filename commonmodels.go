package pokeapi

type APIResource struct {
	Url string `json:"url,omitempty"`
}

type Description struct {
	Description string           `json:"description,omitempty"`
	Language    NamedAPIResource `json:"language,omitempty"`
}

type Effect struct {
	Effect   string           `json:"effect,omitempty"`
	Language NamedAPIResource `json:"language,omitempty"`
}

type Encounter struct {
	MinLevel        int                `json:"min_level,omitempty"`
	MaxLevel        int                `json:"max_level,omitempty"`
	ConditionValues []NamedAPIResource `json:"condition_values,omitempty"`
	Chance          int                `json:"chance,omitempty"`
	Method          NamedAPIResource   `json:"method,omitempty"`
}

type FlavorText struct {
	FlavorText string           `json:"flavor_text,omitempty"`
	Language   NamedAPIResource `json:"language,omitempty"`
	Version    NamedAPIResource `json:"version,omitempty"`
}

type GenerationGameIndex struct {
	GameIndex  int              `json:"game_index,omitempty"`
	Generation NamedAPIResource `json:"geneartion,omitempty"`
}

type MachineVersionDetail struct {
	Machine      APIResource      `json:"machine,omitempty"`
	VersionGroup NamedAPIResource `json:"version_group,omitempty"`
}

type Name struct {
	Name     string           `json:"name,omitempty"`
	Language NamedAPIResource `json:"Language,omitempty"`
}

type NamedAPIResource struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type VerboseEffect struct {
	Effect      string           `json:"effect,omitempty"`
	ShortEffect string           `json:"short_effec,omitemptyt"`
	Language    NamedAPIResource `json:"language,omitempty"`
}

type VersionEncounterDetail struct {
	Version          NamedAPIResource `json:"version,omitempty"`
	MaxChance        int              `json:"max_chance,omitempty"`
	EncounterDetails []Encounter      `json:"encounter_details,omitempty"`
}

type VersionGameIndex struct {
	GameIndex int              `json:"game_index,omitempty"`
	Version   NamedAPIResource `json:"version,omitempty"`
}

type VersionGroupFlavorText struct {
	Text         string           `json:"text,omitempty"`
	Language     NamedAPIResource `json:"language,omitempty"`
	VersionGroup NamedAPIResource `json:"version_group,omitempty"`
}
