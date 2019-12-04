package handler

import (
	"github.com/wesovilabs/beyond-examples/smartcache/storage"
	"net/http"
)

func ListCountries(w http.ResponseWriter, r *http.Request, db storage.Database) {
	if countries, err := db.ListCountries(); err != nil {
		RespondWithError(w, http.StatusNotFound, err)
	} else {
		RespondWithJSON(w, http.StatusOK, countries)
	}
}
