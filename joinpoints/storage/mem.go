package storage

import (
	"errors"
	"fmt"
	"github.com/wesovilabs/goaexamples/joinpointshelper"
	"github.com/wesovilabs/goaexamples/joinpointsmodel"
)

type memDBClient struct {
	data map[string]*model.Employee
}

func newMemDatabase() *memDBClient {
	return &memDBClient{
		data: make(map[string]*model.Employee),
	}
}

func (db *memDBClient) SaveEmployee(employee *model.Employee) error {
	if employee.Email==""{
		return errors.New("invalid email")
	}
	employee.ID = helper.RandomString(10)
	db.data[employee.ID] = employee
	return nil
}

func (db *memDBClient) DeleteEmployee(employeeID string) error {
	if _, ok := db.data[employeeID]; !ok {
		return errors.New(fmt.Sprintf("employee with id %s not found", employeeID))
	}
	delete(db.data, employeeID)
	return nil
}

func (db *memDBClient) GetEmployee(employeeID string) (*model.Employee, error) {
	if employee, ok := db.data[employeeID]; !ok {
		return nil, errors.New(fmt.Sprintf("employee with id %s not found", employeeID))
	} else {
		return employee, nil
	}
}

func (db *memDBClient) ListEmployees() ([]*model.Employee, error) {
	employees := make([]*model.Employee, 0)
	for _, employee := range db.data {
		employees = append(employees, employee)
	}
	return employees, nil
}
