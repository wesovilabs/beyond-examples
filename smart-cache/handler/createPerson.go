package handler

import (
	"encoding/json"
	"github.com/wesovilabs/beyond-examples/smartcache/model"
	"github.com/wesovilabs/beyond-examples/smartcache/storage"
	"net/http"
)

func CreatePerson(w http.ResponseWriter, r *http.Request, db storage.Database) {
	person := model.Person{}
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		RespondWithError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()
	if err := db.AddPerson(&person); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err)
		return
	}
	RespondWithJSON(w, http.StatusCreated, person)
}
