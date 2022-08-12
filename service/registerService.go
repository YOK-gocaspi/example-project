package service

import (
	"errors"
	"example-project/model"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . DatabaseInterface
type DatabaseInterface interface {
	UpdateMany(docs []interface{}) interface{}
	GetByID(id string) model.Employee
	DeleteByID(id string) (int64, error)
	GetAll() []model.Employee
}

type EmployeeService struct {
	DbService DatabaseInterface
}

func NewEmployeeService(dbInterface DatabaseInterface) EmployeeService {
	return EmployeeService{
		DbService: dbInterface,
	}
}

func (s EmployeeService) CreateEmployees(employees []model.Employee) interface{} {

	var emp []interface{}
	for _, employee := range employees {
		emp = append(emp, employee)

	}
	return s.DbService.UpdateMany(emp)
}

func (s EmployeeService) DeleteEmployees(ids []string) (interface{}, error) {
	var deletedIDs []string

	for _, id := range ids {
		deletedCount, err := s.DbService.DeleteByID(id)
		if err != nil {
			return deletedIDs, err
		}

		if deletedCount == 0 {
			err = errors.New("no user has been deleted")
			return deletedIDs, err
		} else {
			deletedIDs = append(deletedIDs, id)
		}
	}

	return deletedIDs, nil
}

func (s EmployeeService) GetEmployeeById(id string) model.Employee {
	return s.DbService.GetByID(id)
}

func (s EmployeeService) GetAllEmployees() []model.Employee {
	return s.DbService.GetAll()
}
