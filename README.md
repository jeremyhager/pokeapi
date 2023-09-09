# Installation
Install:  
```shell
go get github.com/jeremyhager/pokeapi
```

Use:  
```Go
import "github.com/jeremyhager/pokeapi"
```

# About
An SDK/wrapper for pokeapi v2 in Go, baesd on the SDK linked from pokeapi.co https://github.com/mtslzr/pokeapi-go/

## Improvements

This SDK has some code improvements over `pokeapi-go` as well as updated types. For example, `Language` can contain a nested array of the `Name` struct. Within the original `pokeapi-go` package this is defined each time, while this SDK will define this once:

```Go
// github.com/mtslzr/pokeapi-go/blob/master/structs/utility.go
type EvolutionChain struct {
	BabyTriggerItem interface{} `json:"baby_trigger_item"`
	Chain           struct {
		EvolutionDetails []interface{} `json:"evolution_details"`
		EvolvesTo        []struct {
			EvolutionDetails []struct {
    // ...
```

```Go
// github.com/jeremyhager/pokeapi/utility.go
type EvolutionChain struct {
	ID              int               `json:"id"`
	BabyTriggerItem *NamedAPIResource `json:"baby_trigger_item,omitempty"`
	Chain           ChainLink         `json:"chain"`
}

type ChainLink struct {
	IsBaby           bool              `json:"is_baby,omitempty"`
	Species          NamedAPIResource  `json:"species"`
	EvolutionDetails []EvolutionDetail `json:"evolution_details,omitempty"`
	EvolvesTo        []ChainLink       `json:"evolves_to,omitempty"`
}
```


