package pokeapi

type Language struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Official bool   `json:"official"`
	ISO639   string `json:"iso639"`
	ISO3166  string `json:"iso3166"`
	Names    []Name `json:"names"`
}

func GetLanguage(id string) (Language, error) {
	var language Language
	pokeapiUrl := client.urlBuilder(endpoints["language"], id)
	err := client.Get(pokeapiUrl, &language)
	return language, err
}
