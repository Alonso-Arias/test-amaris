package model

type Pokemon struct {
	ID           int32
	Type         string
	Ability      string
	Strength     string
	Moves        string
	Name         string
	VersionGroup string
	IsBattleOnly bool
	IsDefault    bool
	IsMega       bool
}

func (Pokemon) TableName() string {
	return "POKEMONS"
}
