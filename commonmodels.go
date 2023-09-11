package pokeapi

type Languages struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Official bool   `json:"official"`
	ISO639   string `json:"iso639"`
	ISO3166  string `json:"iso3166"`
	Names    []Name `json:"names"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Description struct {
	Description string           `json:"description,omitempty"`
	Language    NamedAPIResource `json:"language,omitempty"`
}

type VersionGameIndex struct {
	GameIndex int              `json:"game_index,omitempty"`
	Version   NamedAPIResource `json:"version,omitempty"`
}

type Name struct {
	Name     string           `json:"name,omitempty"`
	Language NamedAPIResource `json:"Language,omitempty"`
}

type APIResource struct {
	Url string `json:"url,omitempty"`
}

type FlavorText struct {
	FlavorText string           `json:"flavor_text,omitempty"`
	Language   NamedAPIResource `json:"language,omitempty"`
	Version    NamedAPIResource `json:"version,omitempty"`
}
