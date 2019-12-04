package main

import (
	"fmt"
	"github.com/wesovilabs/beyond-examples/smartcache/model"
	"github.com/wesovilabs/beyond-examples/smartcache/test/client"
	"log"
	"net/http"
	"net/url"
	"time"
)

func setUpApiClient() *client.Api {
	u, err := url.Parse("http://localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	return &client.Api{
		BaseURL:    u,
		HTTPClient: http.DefaultClient,
	}
}

var api = setUpApiClient()

func listPeople(cityID string) {
	people, _ := api.ListPeopleByCity(cityID)
	fmt.Printf("There're %v people living in %s\n", len(people), cityID)
}

func main() {
	fmt.Printf("Adding person that lives in MAD\n")
	api.CreatePerson(&model.Person{
		Fullname: "John",
		Email:    "john@mail.com",
		City:     "MAD",
	})
	listPeople("MAD")
	fmt.Printf("Adding person that lives in MAD\n")
	api.CreatePerson(&model.Person{
		Fullname: "Jane",
		Email:    "jane@mail.com",
		City:     "MAD",
	})
	fmt.Printf("Adding person that lives in MLG\n")
	api.CreatePerson(&model.Person{
		Fullname: "Marina",
		Email:    "marina@mail.com",
		City:     "MLG",
	})
	listPeople("MAD")
	listPeople("MLG")
	fmt.Println("Sleeping 10 seconds...")
	time.Sleep(10 * time.Second)
	listPeople("MLG")
	listPeople("MAD")
}
