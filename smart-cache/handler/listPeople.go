package handler

import (
	"github.com/go-chi/chi"
	"github.com/wesovilabs/beyond-examples/smartcache/storage"
	"net/http"
)

func ListPeople(w http.ResponseWriter, r *http.Request, db storage.Database) {
	cityID := chi.URLParam(r, "cityId")
	if people, err := db.ListPeopleByCity(cityID); err != nil {
		RespondWithError(w, http.StatusNotFound, err)
	} else {
		RespondWithJSON(w, http.StatusOK, people)
	}
}
