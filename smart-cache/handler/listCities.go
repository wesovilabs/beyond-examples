package handler

import (
	"github.com/go-chi/chi"
	"github.com/wesovilabs/beyond-examples/smartcache/storage"
	"net/http"
)

func ListCities(w http.ResponseWriter, r *http.Request, db storage.Database) {
	countryID := chi.URLParam(r, "countryId")
	if cities, err := db.ListCitiesByCountry(countryID); err != nil {
		RespondWithError(w, http.StatusNotFound, err)
	} else {
		RespondWithJSON(w, http.StatusOK, cities)
	}
}
