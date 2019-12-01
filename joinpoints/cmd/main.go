package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/wesovilabs/beyond/api"
	"github.com/wesovilabs/beyond-examples/joinpoints/advice"
	"github.com/wesovilabs/beyond-examples/joinpoints/handler"
	"github.com/wesovilabs/beyond-examples/joinpoints/storage"
	"net/http"
)


func Beyond() *api.Beyond {
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
