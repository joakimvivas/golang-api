package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type API struct{}

type ContactsParams struct {
	Offset int `schema:"offset"`
	Limit  int `schema:"limit"`
}

type PostContact struct {
	Name string `json:"name"`
}

var (
	contacts = []string{"Contact 1", "Contact 2", "Contact 3"}
	decoder  = schema.NewDecoder() // convert schema
)

func (a *API) getContacts(w http.ResponseWriter, r *http.Request) {
	params := &ContactsParams{}

	err := decoder.Decode(params, r.URL.Query())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if params.Offset > len(contacts) || params.Offset < 0 { //outside limits
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if params.Limit < 0 || params.Limit > len(contacts) { //outside offset
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var from, to int
	if params.Offset > 0 {
		from = params.Offset
	}

	if params.Limit > 0 {
		to = params.Limit
	} else {
		to = len(contacts)
	}

	json.NewEncoder(w).Encode(contacts[from:to]) //all values of contacts from params.offset until params.limit
}

func (a *API) getContact(w http.ResponseWriter, r *http.Request) {
	pathsParams := mux.Vars(r)

	idParam := pathsParams["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	index := id - 1

	if index < 0 || index > len(contacts)-1 { //contacts-1
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(contacts[index])

}

func (a *API) postContact(w http.ResponseWriter, r *http.Request) {

	contact := &PostContact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contacts = append(contacts, contact.Name)
	w.WriteHeader(http.StatusCreated)
}
