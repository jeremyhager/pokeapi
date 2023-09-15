package pokeapi

import "net/url"

var (
	papiDefault = url.URL{
		Scheme: "https",
		Host:   "pokeapi.co",
		Path:   "/api/v2",
	}
	client  *RESTClient
	papiEnv = "POKEAPI_URL"
)

// endpoints map from key:string to value:string. To change the endpoint, change the value.
//
// Example, chaging the ability endpoint:
//
//	"ability": "ability", change to "ability": "someOtherEndpoint"
var endpoints = map[string]string{
	"ability":                   "ability",
	"berry":                     "berry",
	"berry-firmness":            "berry-firmness",
	"berry-flavor":              "berry-flavor",
	"characteristic":            "characteristic",
	"contest-effect":            "contest-effect",
	"contest-type":              "contest-type",
	"egg-group":                 "egg-group",
	"encounter-condition":       "encounter-condition",
	"encounter-condition-value": "encounter-condition-value",
	"encounter-method":          "encounter-method",
	"evolution-chain":           "evolution-chain",
	"evolution-trigger":         "evolution-trigger",
	"gender":                    "gender",
	"generation":                "generation",
	"growth-rate":               "growth-rate",
	"item":                      "item",
	"item-attribute":            "item-attribute",
	"item-category":             "item-category",
	"item-fling-effect":         "item-fling-effect",
	"item-pocket":               "item-pocket",
	"language":                  "language",
	"location":                  "location",
	"location-area":             "location-area",
	"machine":                   "machine",
	"move":                      "move",
	"move-ailment":              "move-ailment",
	"move-battle-style":         "move-battle-style",
	"move-category":             "move-category",
	"move-damage-class":         "move-damage-class",
	"move-learn-method":         "move-learn-method",
	"move-target":               "move-target",
	"nature":                    "nature",
	"pal-park-area":             "pal-park-area",
	"pokeathlon-stat":           "pokeathlon-stat",
	"pokedex":                   "pokedex",
	"pokemon":                   "pokemon",
	"pokemon-color":             "pokemon-color",
	"pokemon-form":              "pokemon-form",
	"pokemon-habitat":           "pokemon-habitat",
	"pokemon-shape":             "pokemon-shape",
	"pokemon-species":           "pokemon-species",
	"region":                    "region",
	"stat":                      "stat",
	"super-contest-effect":      "super-contest-effect",
	"type":                      "type",
	"version":                   "version",
	"version-group":             "version-group",
}
