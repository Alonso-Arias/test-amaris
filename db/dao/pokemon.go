package dao

import (
	"context"

	"github.com/Alonso-Arias/test-amaris/db/base"
	"github.com/Alonso-Arias/test-amaris/db/model"
	"github.com/Alonso-Arias/test-amaris/log"
	"gorm.io/gorm"
)

var loggerf = log.LoggerJSON().WithField("package", "dao")

// PokemonDAO - Pokemon dao interface
type PokemonDAO interface {
	Get(ctx context.Context, sku int32) (model.Pokemon, error)
	Save(ctx context.Context, Pokemon model.Pokemon) error
}

// PokemonDAOImpl - Pokemon dao implementation
type PokemonDAOImpl struct {
}

// NewPokemonDAO - gets an PokemonDAOImpl instance
func NewPokemonDAO() *PokemonDAOImpl {
	return &PokemonDAOImpl{}
}

// FindAll -
func (pd *PokemonDAOImpl) Get(ctx context.Context, sku int32) (model.Pokemon, error) {

	log := loggerf.WithField("struct", "PokemonDAOImpl").WithField("function", "Get")

	db := base.GetDB()

	Pokemon := model.Pokemon{}
	err := db.Where("id = ?", sku).FirstOrInit(&Pokemon).Error

	if err != nil {
		log.WithError(err).Error("get Pokemons fails")
		return model.Pokemon{}, err
	} else if Pokemon.ID == 0 {
		return model.Pokemon{}, gorm.ErrRecordNotFound
	}

	log.Debugf("%v", Pokemon)

	return Pokemon, nil

}

func (pd *PokemonDAOImpl) Save(ctx context.Context, Pokemon model.Pokemon) error {

	log := loggerf.WithField("struct", "PokemonDAOImpl").WithField("function", "Save")

	db := base.GetDB()

	err := db.Create(&Pokemon)

	if err.Error != nil {
		log.Debugf("%v", err.Error)
		return err.Error
	}

	log.Infof("Save Pokemon Sucessfull\n")

	return nil

}
