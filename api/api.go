package api

import (
	"encoding/json"
	"net/http"
)

// Coin Balance Request Params
type CoinBalanceRequestParams struct {
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	// Success HTTP code
	Code int

	// Account Balance
	Balance int64
}

// Error Response
type Error struct {
	// Error HTTP code
	Code int

	// Error message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {

	// create instance of Error struct for response
	resp := Error{
		Code:    code,
		Message: message,
	}

	// set the content type header, as we want to return JSON
	w.Header().Set("Content-Type", "application/json")
	// set the error code
	w.WriteHeader(code)

	// write error struct out as JSON back to caller
	json.NewEncoder(w).Encode(resp)
}

// set up variables to handle different error scenarios using writeError
var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An unexpected error occurred.", http.StatusInternalServerError)
	}
)
