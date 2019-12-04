package storage

import (
	"errors"
	"fmt"
	"github.com/wesovilabs/beyond-examples/smartcache/model"
)

type memDBClient struct {
	people    map[string]*model.Person
	countries map[string]*model.Country
	cities    map[string]*model.City
}

func newMemDatabase() *memDBClient {
	return &memDBClient{
		people:    make(map[string]*model.Person),
		countries: make(map[string]*model.Country),
		cities:    make(map[string]*model.City),
	}
}

func (db *memDBClient) AddPerson(person *model.Person) error {
	if person.Email == "" {
		return errors.New("invalid email")
	}
	person.ID = generateID(10)
	db.people[person.ID] = person
	return nil
}

func (db *memDBClient) DeletePerson(personID string) error {
	if _, ok := db.people[personID]; !ok {
		return errors.New(fmt.Sprintf("person with id %s not found", personID))
	}
	delete(db.people, personID)
	return nil
}

func (db *memDBClient) ListCountries() ([]*model.Country, error) {
	countries := make([]*model.Country, 0)
	for _, country := range db.countries {
		countries = append(countries, country)
	}
	return countries, nil
}

func (db *memDBClient) ListCitiesByCountry(countryID string) ([]*model.City, error) {
	cities := make([]*model.City, 0)
	for _, city := range db.cities {
		if city.Country == countryID {
			cities = append(cities, city)
		}
	}
	return cities, nil
}

func (db *memDBClient) ListPeopleByCity(cityID string) ([]*model.Person, error) {
	fmt.Println("accessing to the database...")
	people := make([]*model.Person, 0)
	for _, person := range db.people {
		if person.City == cityID {
			people = append(people, person)
		}
	}
	return people, nil
}

func (db *memDBClient) SetUp() {
	for _,country:=range countries {
		db.countries[country.ID]=country
	}
	for _,city:=range cities {
		db.cities[city.ID]=city
	}
}

