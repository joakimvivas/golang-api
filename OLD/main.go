package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joakimvivas/golang-api/api"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{\"message\": \"Hello World\"}")
}

func main() {

	r := mux.NewRouter()

	// MongoDB Host and Port
	host := "localhost"
	port := 27017

	// create the api object
	a := &api.API{}

	// register the routes
	a.RegisterRoutes(r)
	r.HandleFunc("/", handleIndex).Methods(http.MethodGet)

	// MongoDB connection
	log.Println("Starting the MongoDB connection...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port))
	client, error := mongo.Connect(ctx, clientOpts)
	if error != nil {
		panic(error)
	}

	// Check the MongoDB connections
	error = client.Ping(context.TODO(), nil)
	if error != nil {
		panic(error)
	}

	// Starting Golang server
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
