package handler

import (
	"github.com/go-chi/chi"
	"github.com/wesovilabs/beyond-examples/smartcache/storage"
	"net/http"
)

func DeletePerson(w http.ResponseWriter, r *http.Request, db storage.Database) {
	id := chi.URLParam(r, "id")
	if err := db.DeletePerson(id); err != nil {
		RespondWithError(w, http.StatusNotFound, err)
		return
	}
	RespondWithJSON(w, http.StatusNoContent, nil)

}
