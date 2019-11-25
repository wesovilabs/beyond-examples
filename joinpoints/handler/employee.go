package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/wesovilabs/goaexamples/joinpointshandler/internal"
	"github.com/wesovilabs/goaexamples/joinpointsmodel"
	"github.com/wesovilabs/goaexamples/joinpointsstorage"
	"net/http"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request, db storage.Database) {
	req := &EmployeeRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		internal.RespondWithError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()
	employee := &model.Employee{
		Fullname: req.Fullname,
		Email:    req.Email,
	}
	if err := db.SaveEmployee(employee); err != nil {
		internal.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}
	internal.RespondWithJSON(w, http.StatusCreated,employee)
}

func ListEmployees(w http.ResponseWriter, r *http.Request, db storage.Database) {
	if employees, err := db.ListEmployees(); err != nil {
		internal.RespondWithError(w, http.StatusNotFound, err)
	} else {
		internal.RespondWithJSON(w, http.StatusOK, employees)
	}
}

func GetEmployee(w http.ResponseWriter, r *http.Request, db storage.Database) {
	id := chi.URLParam(r, "id")
	if employee, err := db.GetEmployee(id); err != nil {
		internal.RespondWithError(w, http.StatusNotFound, err)
	} else {
		internal.RespondWithJSON(w, http.StatusOK, employee)
	}
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request, db storage.Database) {
	id := chi.URLParam(r, "id")
	if err:=db.DeleteEmployee(id); err != nil {
		internal.RespondWithError(w, http.StatusNotFound, err)
		return
	}
	internal.RespondWithJSON(w, http.StatusNoContent, nil)

}

type EmployeeRequest struct {
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}

type EmployeeResponse struct {
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}
