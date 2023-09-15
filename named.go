package pokeapi

type Named struct {
	Count    int                `json:"count"`
	Next     string             `json:"next"`
	Previous string             `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

func GetNamedEndpoint(endpoint string) (Named, error) {
	var named Named
	pokeapiUrl := client.urlBuilder(endpoints[endpoint], "")
	err := client.Get(pokeapiUrl, &named)
	return named, err
}
