package dao

import (
	"context"
	"testing"

	"github.com/Alonso-Arias/test-amaris/db/model"
	"github.com/stretchr/testify/assert"
)

var pokemonDao = NewPokemonDAO()

func TestSave_OK(t *testing.T) {

	err := pokemonDao.Save(context.TODO(), model.Pokemon{ID: 32, Type: "", Ability: "", Strength: "", Moves: ""})

	if err != nil {
		assert.FailNowf(t, "fails", "fails to saveing Pokemon: %v", err)
	}

}
