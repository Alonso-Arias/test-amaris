package pokemon

import (
	"context"
	"fmt"

	"github.com/Alonso-Arias/test-amaris/db/dao"
	md "github.com/Alonso-Arias/test-amaris/db/model"
	errs "github.com/Alonso-Arias/test-amaris/errors"
	ext "github.com/Alonso-Arias/test-amaris/external"
	"github.com/Alonso-Arias/test-amaris/log"
	"github.com/Alonso-Arias/test-amaris/services/model"
	"gopkg.in/dealancer/validate.v2"
	"gorm.io/gorm"
)

var loggerf = log.LoggerJSON().WithField("package", "services")

type PokemonService struct {
}

type SavePokemonRequest struct {
	Pokemon model.Pokemon `json:"pokemon"`
}
type SavePokemonResponse struct {
}

func (ps PokemonService) SavePokemon(ctx context.Context, in SavePokemonRequest) (SavePokemonResponse, error) {
	log := loggerf.WithField("service", "PokemonService").WithField("func", "GetPokemon")

	log.Info("start")
	defer log.Info("finish")

	// validates input request
	if err := validate.Validate(in); err != nil {
		log.WithError(err).Error("validates problems")
		return SavePokemonResponse{}, errs.BadRequest
	}

	// Obtener información externa del Pokemon desde la API
	extPokemon, err := ext.GetExternalPokemon(in.Pokemon.ID)
	if err != nil {
		return SavePokemonResponse{}, errs.PokemonsNotFound
	}

	PokemonDao := dao.NewPokemonDAO()

	// validacion de Pokemon
	p, err := PokemonDao.Get(ctx, int32(in.Pokemon.ID))
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			log.WithError(err).Error("problems with getting Pokemon")
			return SavePokemonResponse{}, err
		}
	}

	// valida si ya esta o no el Pokemon a guardar
	if p.ID == int32(in.Pokemon.ID) {
		return SavePokemonResponse{}, errs.PokemonAlreadySaved
	}

	// decision con respecto al nombre dejarlo dinamico con interface
	var namePokemon string
	switch p := extPokemon.Pokemon.(type) {
	case string:
		// Manejar el caso en el que Pokemon es un string
		namePokemon = p
		fmt.Println("Nombre del Pokémon:", p)
	case map[string]interface{}:
		// Manejar el caso en el que Pokemon es un objeto
		nombre, ok := p["name"].(string)
		if ok {
			log.Debugf("Nombre del Pokemon:", nombre)
			namePokemon = nombre
		} else {
			log.Errorf("No se pudo obtener el nombre del Pokémon.")
		}
	default:
		fmt.Println("Tipo de Pokemon no reconocido.")
	}

	err = PokemonDao.Save(ctx, md.Pokemon(
		md.Pokemon{
			ID:           int32(in.Pokemon.ID),
			Type:         in.Pokemon.Type,
			Ability:      in.Pokemon.Ability,
			Strength:     in.Pokemon.Strength,
			Moves:        in.Pokemon.Moves,
			Name:         namePokemon,
			VersionGroup: extPokemon.VersionGroup.Name,
			IsBattleOnly: extPokemon.IsBattleOnly,
			IsDefault:    extPokemon.IsDefault,
			IsMega:       extPokemon.IsMega,
		}))

	if err != nil {
		return SavePokemonResponse{}, err
	}

	return SavePokemonResponse{}, nil
}
