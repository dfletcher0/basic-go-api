package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dfletcher0/basic-go-api/api"
	"github.com/dfletcher0/basic-go-api/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceRequestParams{}

	// decode request qparams into CoinBalanceRequestParams struct:
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	// grab username from URL & put into struct
	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		// then there was an error decoding the query params
		// raise internal error
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		// then there was an error creating the database
		// raise internal error
		api.InternalErrorHandler(w)
		return
	}

	// now we have a database, get the required information
	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// set response struct
	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code:    http.StatusOK,
	}

	// return response to caller
	w.Header().Set("Content-Type", "application.json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
