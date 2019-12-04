// bui
package storage

import (
	"github.com/wesovilabs/beyond-examples/smartcache/model"
)

type Database interface {
	AddPerson(person *model.Person) error
	DeletePerson(personID string) error
	ListCountries() ([]*model.Country, error)
	ListCitiesByCountry(countryID string) ([]*model.City, error)
	ListPeopleByCity(cityID string) ([]*model.Person, error)
}

func New() Database {
	db:=newMemDatabase()
	db.SetUp()
	return db
}
