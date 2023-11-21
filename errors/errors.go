package errors

import (
	"fmt"
)

// CustomError custom error
type CustomError struct {
	Message      string `json:"message"`
	Code         int    `json:"code"`
	InternalCode string `json:"internalCode"`
}

func (e CustomError) SetMessage(msg string) CustomError {
	e.Message = msg
	return e
}
func (e CustomError) Error() string {
	return fmt.Sprintf("error %d: %s", e.Code, e.Message)
}

var (
	BadRequest    = CustomError{Message: "BadRequest", Code: 400, InternalCode: "BADREQUEST"}
	Unauthorized  = CustomError{Message: "Unauthorized", Code: 401, InternalCode: "UNAUTHORIZED"}
	NotFound      = CustomError{Message: "NotFound", Code: 404, InternalCode: "NOT_FOUND"}
	InternalError = CustomError{Message: "Error", Code: 500, InternalCode: "INTERNAL_SERVER_ERROR"}

	PokemonsNotFound    = CustomError{Message: "Pokemons not found", Code: 404, InternalCode: "POKEMONS_NOT_FOUND"}
	PokemonAlreadySaved = CustomError{Message: "Pokemon already saved", Code: 400, InternalCode: "POKEMON_ALREADY_SAVED"}
)
