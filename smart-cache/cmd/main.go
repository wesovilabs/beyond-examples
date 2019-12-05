package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/wesovilabs/beyond-examples/smartcache/advice"
	"github.com/wesovilabs/beyond-examples/smartcache/handler"
	"github.com/wesovilabs/beyond-examples/smartcache/storage"
	"github.com/wesovilabs/beyond/api"
	"net/http"
)

func Beyond()*api.Beyond{
	return api.New().
		WithAround(advice.
			Memoize(5),
			"storage.*memDBClient.ListPeopleByCity(...)...")
}

func main() {
	storage := storage.New()
	r := chi.NewRouter()
	r.Post("/people", func(w http.ResponseWriter, r *http.Request) {
		handler.CreatePerson(w, r, storage)
	})
	r.Delete("/people/{id}", func(w http.ResponseWriter, r *http.Request) {
		handler.DeletePerson(w, r, storage)
	})
	r.Get("/countries", func(w http.ResponseWriter, r *http.Request) {
		handler.ListCountries(w, r, storage)
	})
	r.Get("/countries/{countryId}/cities", func(w http.ResponseWriter, r *http.Request) {
		handler.ListCities(w, r, storage)
	})
	r.Get("/cities/{cityId}/people", func(w http.ResponseWriter, r *http.Request) {
		handler.ListPeople(w, r, storage)
	})
	fmt.Println("Launching server on localhost:8000")
	if err := http.ListenAndServe("0.0.0.0:8000", r); err != nil {
		panic(err)
	}

}
