package pokeapi

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Cost        int    `json:"cost"`
	FlingPower  int    `json:"fling_power"`
	FlingEffect struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"fling_effect"`
	Attributes []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"attributes"`
	Category struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"category"`
	EffectEntries []struct {
		Effect      string `json:"effect"`
		ShortEffect string `json:"short_effect"`
		Language    struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	}
	FlavorText []struct {
		Text         string `json:"text"`
		VersionGroup struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version_group"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	}
	GameIndices []struct {
		GameIndex  int `json:"game_index"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
	}
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	}
	Sprites struct {
		Default string `json:"default"`
	} `json:"sprites"`
	HeldByPokemon []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"held_by_pokemon"`
	BabyTriggerFor struct {
		URL string `json:"url"`
	} `json:"baby_trigger_for"`
}
