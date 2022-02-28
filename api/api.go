package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type API struct{}

var contacts = []string{"Contact 1", "Contact 2", "Contact 3"}

func (a *API) getContacts(w http.ResponseWriter, r *http.Request) {
	limitParam := r.URL.Query().Get("limit") //limit of value request
	limit, err := strconv.Atoi(limitParam)   //convert to int
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if limit < 0 || limit > len(contacts) { //outside limits
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(contacts[0:limit]) //all values of contacts until limit
}

func (a *API) getContact(w http.ResponseWriter, r *http.Request) {}
