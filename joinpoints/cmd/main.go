package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/wesovilabs/goa/api"
	"github.com/wesovilabs/goaexamples/joinpointsadvice"
	"github.com/wesovilabs/goaexamples/joinpointshandler"
	"github.com/wesovilabs/goaexamples/joinpointsstorage"
	"net/http"
)


func Goa() *api.Goa {
	return api.New().
		WithBefore(advice.NewSimpleTracingAdvice, `*.*(...)...`)
}

func main() {
	storage := storage.New()
	r := chi.NewRouter()
	r.Post("/employees", func(w http.ResponseWriter, r *http.Request) {
		handler.CreateEmployee(w, r, storage)
	})
	r.Get("/employees", func(w http.ResponseWriter, r *http.Request) {
		handler.ListEmployees(w, r, storage)
	})
	r.Get("/employees/{id}", func(w http.ResponseWriter, r *http.Request) {
		handler.GetEmployee(w, r, storage)
	})
	r.Delete("/employees/{id}", func(w http.ResponseWriter, r *http.Request) {
		handler.DeleteEmployee(w, r, storage)
	})
	fmt.Println("Launching server on localhost:8000")
	if err := http.ListenAndServe("0.0.0.0:8000", r); err != nil {
		panic(err)
	}
}
