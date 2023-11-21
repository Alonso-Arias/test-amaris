package model

// swagger:model Pokemon
type Pokemon struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Ability  string `json:"ability"`
	Strength string `json:"strength"`
	Moves    string `json:"moves"`
}

type PokemonExternal struct {
	ID           int         `json:"id"`
	IsBattleOnly bool        `json:"is_battle_only"`
	IsDefault    bool        `json:"is_default"`
	IsMega       bool        `json:"is_mega"`
	Pokemon      interface{} `json:"pokemon"`
	VersionGroup struct {
		Name string `json:"name"`
	} `json:"version_group"`
}
