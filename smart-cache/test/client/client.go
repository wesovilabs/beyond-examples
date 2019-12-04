package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/wesovilabs/beyond-examples/smartcache/model"
	"io"
	"net/http"
	"net/url"
)

type Api struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
}

func (c *Api) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func (c *Api) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

func (c *Api) CreatePerson(person *model.Person) (*model.Person, error) {
	req, err := c.newRequest(http.MethodPost, "/people", person)
	if err != nil {
		return nil, err
	}
	output := &model.Person{}
	_, err = c.do(req, output)
	return output, err
}

func (c *Api) DeletePerson(personID string) error {
	req, err := c.newRequest(http.MethodDelete, "/people/"+personID, nil)
	if err != nil {
		return err
	}
	_, err = c.do(req, nil)
	return err
}

func (c *Api) ListCountries() ([]*model.Country, error) {
	req, err := c.newRequest(http.MethodGet, "/countries", nil)
	if err != nil {

		return nil, err
	}
	var output []*model.Country
	_, err = c.do(req, &output)
	return output, err
}

func (c *Api) ListCitiesByCountry(countryID string) ([]*model.City, error) {
	path := fmt.Sprintf("/countries/%s/cities", countryID)
	req, err := c.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var output []*model.City
	_, err = c.do(req, &output)
	return output, err
}

func (c *Api) ListPeopleByCity(cityID string) ([]*model.Person, error) {
	path := fmt.Sprintf("/cities/%s/people", cityID)
	req, err := c.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var output []*model.Person
	_, err = c.do(req, &output)
	return output, err
}
