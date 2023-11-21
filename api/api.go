package main

import (
	"context"
	"net/http"

	_ "github.com/Alonso-Arias/test-amaris/api/docs"
	errs "github.com/Alonso-Arias/test-amaris/errors"
	"github.com/Alonso-Arias/test-amaris/log"
	"github.com/Alonso-Arias/test-amaris/services/pokemon"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var loggerf = log.LoggerJSON().WithField("package", "main")

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1323
// @BasePath /api/v1
func main() {
	e := echo.New()
	e.POST("/api/v1/pokemon", pokemonPost)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":1323"))

}

// save pokemon
// @Summary save pokemon
// @tags pokemon
// @Description guarda un pokemon
// @ID pokemonPost
// @Accept  json
// @Produce  json
// @Param SavepokemonRequest body pokemon.SavepokemonRequest true "pokemon"
// @Success 200  {object} pokemon.SavePokemonResponse
// @Failure 404 {object}  errors.CustomError
// @Failure 500 {object}  errors.CustomError
// @Router /pokemon [post]
func pokemonPost(c echo.Context) error {

	log := loggerf.WithField("func", "pokemonPost")

	req := &pokemon.SavePokemonRequest{}

	if err := c.Bind(req); err != nil {
		log.WithError(err).Error("Binding error")
		return c.JSON(http.StatusBadRequest, err)
	}

	res, err := pokemon.PokemonService{}.SavePokemon(context.TODO(), *req)
	if ce, ok := err.(errs.CustomError); ok {
		return c.JSON(ce.Code, err)
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}
