package client

import (
	"bytes"
	"encoding/json"
	"github.com/wesovilabs/goaexamples/joinpoints/model"
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

func (c *Api) CreateEmployee(employee *model.Employee) (*model.Employee, error) {
	req, err := c.newRequest(http.MethodPost, "/employees", employee)
	if err != nil {
		return nil, err
	}
	output:=&model.Employee{}
	_, err = c.do(req, output)
	return output, err
}

func (c *Api) GetEmployee(employeeID string) (*model.Employee, error) {
	req, err := c.newRequest(http.MethodGet, "/employees/"+employeeID, nil)
	if err != nil {
		return nil, err
	}
	var output *model.Employee
	_, err = c.do(req, output)
	return output, err
}

func (c *Api) DeleteEmployee(employeeID string) error {
	req, err := c.newRequest(http.MethodDelete, "/employees/"+employeeID, nil)
	if err != nil {
		return err
	}
	_, err = c.do(req, nil)
	return err
}

func (c *Api) ListEmployees() ([]*model.Employee, error) {
	req, err := c.newRequest(http.MethodGet, "/employees", nil)
	if err != nil {
		return nil, err
	}
	var output []*model.Employee
	_, err = c.do(req, output)
	return output, err
}
