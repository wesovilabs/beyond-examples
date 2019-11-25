package main

import (
	"github.com/wesovilabs/goaexamples/joinpointsmodel"
	"github.com/wesovilabs/goaexamples/joinpointstest/client"
	"log"
	"net/http"
	"net/url"
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

func main() {
	api := setUpApiClient()
	res,_ := api.CreateEmployee(&model.Employee{
		Email:    "john.doe@mail.com",
		Fullname: "John Doe",
	})
	if res.ID!="" {
		api.GetEmployee(res.ID)
		api.ListEmployees()
		api.DeleteEmployee(res.ID)
	}
}
