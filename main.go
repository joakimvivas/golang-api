package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joakimvivas/golang-api/api"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{\"message\": \"Hello World\"}")
}

func main() {

	r := mux.NewRouter()

	// create the api object
	a := &api.API{}

	// register the routes
	a.RegisterRoutes(r)

	r.HandleFunc("/", handleIndex).Methods(http.MethodGet)

	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("Listening on 8080...")
	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}

}
