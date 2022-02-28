package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) RegisterRoutes(r *mux.Router) {

	r.HandleFunc("/contacts", a.getContacts).Methods(http.MethodGet)
	r.HandleFunc("/contacts/{id}", a.getContact).Methods(http.MethodGet)

}
