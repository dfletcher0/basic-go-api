package middleware

import (
	"errors"
	"net/http"

	"github.com/dfletcher0/basic-go-api/api"
	"github.com/dfletcher0/basic-go-tutorial/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnauthorizedError = errors.New("Invalid username or token.")

// all middleware functions need to take in & return http.Handler interfaces
func Authorization(next http.Handler) http.Handler {
	// ResponseWriter is used to construct a response to the caller
	// Request contains all incoming request info
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// extract username query param from request
		var username string = r.URL.Query().Get("username")
		// extract auth token from header
		var token = r.Header.Get("Authorization")
		var err error

		// if either username or token is empty, return an error
		if username == "" || token == "" {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			// exit function (auth failed)
			return
		}

		// at this point: authorization has passed, processing can begin
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		// query database using GetUserLoginDetails method
		var loginDetails *tools.loginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		// if login details cannot be found, or retrieved token doesn't match, exit with unauthorized
		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		// call next middleware in queue, or handler (if no more middleware needs to run)
		// in this case, calls the GetCoinBalance handler function
		next.ServeHTTP(w, r)
	})
}
