package handlers

import (
	"github.com/dfletcher0/basic-go-api/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

// capital starting letter: function is allowed to be imported in other packages (public)
// define middleware:
func Handler(r *chi.Mux) {
	// Global middleware
	// always ignore trailing slashes in HTTP calls
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {

		// Middleware for /account route: every inbound request must pass authorization function
		router.Use(middleware.Authorization)

		// Add HTTP GET method on /account path handled by GetCoinBalance function
		router.Get("/coins", GetCoinBalance)
	})
}
