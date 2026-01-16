package pokeapi

type RespShallowLocations struct {
	Count	 int 	 `json:"count"`
	Next 	 *string `json:"next"`
	Previous *string `json:"previous"`
	Results	 []struct {
		Name	string	`json:"name"`
		URL		string	`json:"url"`
	} `json:"results"`
}


type PokemonLocationData struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name		   string `json:"name"`
	BaseExperience int `json:"base_experience"`
}