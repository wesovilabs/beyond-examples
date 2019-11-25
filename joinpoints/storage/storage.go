package storage

import (
	"github.com/wesovilabs/goaexamples/joinpoints/model"
)

type Database interface {
	SaveEmployee(employee *model.Employee) error
	DeleteEmployee(employeeID string) error
	GetEmployee(employeeID string) (*model.Employee, error)
	ListEmployees() ([]*model.Employee, error)
}

func New() Database {
	return newMemDatabase()
}
