package external

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Alonso-Arias/test-amaris/log"
	"github.com/Alonso-Arias/test-amaris/services/model"
)

var loggerf = log.LoggerJSON().WithField("package", "external")

var urlBase = os.Getenv("BASE_URL")

// se rea la request por detras hacia la url para consumir el api
func GetExternalPokemon(id int) (model.PokemonExternal, error) {
	log := loggerf.WithField("func", "GetExternalPokemon")

	idStr := strconv.Itoa(id)

	req, err := http.NewRequest("GET", urlBase+"/"+idStr, nil)
	if err != nil {
		return model.PokemonExternal{}, err
	}
	req.Header.Set("Accept", "application/json")

	client := getClientCofiguration()

	resp, err := client.Do(req)
	if err != nil {
		return model.PokemonExternal{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.PokemonExternal{}, err
	}
	bodyBytes = bytes.TrimPrefix(bodyBytes, []byte("\xef\xbb\xbf")) // evitar problemas de formato utf

	var localsResponse model.PokemonExternal
	// se parsea el json a la structura declarada
	err = json.Unmarshal(bodyBytes, &localsResponse)
	if err != nil {
		return model.PokemonExternal{}, err
	}

	log.Debugf("Body: %s", string(bodyBytes))

	return localsResponse, err

}

// se configuran los deadines de peticiones hacia la api
func getClientCofiguration() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   time.Duration(5) * time.Second,
				KeepAlive: time.Duration(5),
			}).Dial,
			TLSHandshakeTimeout:   5 * time.Second,
			ResponseHeaderTimeout: 5 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	return client
}
