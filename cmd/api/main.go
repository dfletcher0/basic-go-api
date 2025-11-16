package main

import (
	"fmt"
	"net/http"

	"github.com/dfletcher0/basic-go-api/internals/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

// chi package: flexible easy-to-use web dev package
// import package from own module (internal/handlers folder)
// alias log import
// install external packages using $ go mod tidy

func main() {
	// set up logger to print file & line number
	log.SetReportCaller(true)

	// struct used to configure API
	var r *chi.Mux = chi.NewRouter()

	// configure router with endpoint definitions
	handlers.Handler(r)

	fmt.Println("Starting Go API service...")

	// start server & handle err
	// pass default host port & chi handler
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
