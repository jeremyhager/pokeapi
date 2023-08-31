package commonmodels

type NamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Description struct {
	Description string           `json:"description"`
	Language    NamedAPIResource `json:"language"`
}

type VersionGameIndex struct {
	GameIndex int              `json:"game_index"`
	Version   NamedAPIResource `json:"version"`
}

type Name struct {
	Name     string           `json:"name"`
	Language NamedAPIResource `json:"Language"`
}

type APIResource struct {
	Url string `json:"url"`
}

type FlavorText struct {
	FlavorText string           `json:"flavor_text"`
	Language   NamedAPIResource `json:"language"`
	Version    NamedAPIResource `json:"version"`
}
