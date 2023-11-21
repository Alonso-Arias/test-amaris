package base

import (
	"log"
	"time"

	"testing"

	m "github.com/Alonso-Arias/test-amaris/db/model"
)

func TestGetConnection(t *testing.T) {

	dbc := GetDB()

	result := m.Pokemon{}

	dbc.Raw("SELECT * FROM POKEMONS").Scan(&result)

}

func TestGetTime(t *testing.T) {

	loc, _ := time.LoadLocation("Europe/Monaco")
	//set timezone,
	savetrxTime := time.Now().In(loc)

	log.Println("Hora  : ", savetrxTime)

}
