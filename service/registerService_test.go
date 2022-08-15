package service_test

import (
	"errors"
	"example-project/model"
	"example-project/service"
	"example-project/service/servicefakes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetEmployeeById(t *testing.T) {
	fakeDB := &servicefakes.FakeDatabaseInterface{}

	data := model.Employee{
		ID:        "1",
		FirstName: "jon",
		LastName:  "kock",
		Email:     "jon@gmail.com",
	}

	fakeDB.GetByIDReturns(data)

	serviceInstance := service.NewEmployeeService(fakeDB)
	actual := serviceInstance.GetEmployeeById("1")
	assert.Equal(t, data, actual)

}

func TestCreateEmployees(t *testing.T) {
	//here comes your first unit test which should cover the function CreateEmployees
}

func TestUpdateEmployees(t *testing.T) {
	fakeDB := &servicefakes.FakeDatabaseInterface{}

	fakeDBError := errors.New("fake DB Error")
	// always fails for Employee with ID 42
	// otherwise never error
	fakeDB.UpdateEmployeeCalls(func(employee model.Employee) error {
		if employee.ID == "42" {
			return fakeDBError
		}
		return nil
	})

	var tests = []struct {
		employees      []model.Employee
		expectedResult []string
		expectedErr    error
	}{
		{[]model.Employee{{ID: "1"}}, []string{"1"}, nil},
		{[]model.Employee{{ID: "1"}, {ID: "2"}}, []string{"1", "2"}, nil},
		{[]model.Employee{{ID: "42"}}, []string(nil), fakeDBError},
		{[]model.Employee{{ID: "1"}, {ID: "42"}}, []string{"1"}, fakeDBError},
	}

	for _, tt := range tests {
		serviceInstance := service.NewEmployeeService(fakeDB)
		actualResult, actualErr := serviceInstance.UpdateEmployees(tt.employees)

		assert.Equal(t, tt.expectedErr, actualErr)
		assert.Equal(t, tt.expectedResult, actualResult)
	}
}

func TestGetAllEmployees(t *testing.T) {
	fakeDB := &servicefakes.FakeDatabaseInterface{}
	data := []model.Employee{
		{
			ID:        "1",
			FirstName: "jon",
			LastName:  "kock",
			Email:     "jon@gmail.com",
		},
	}
	fakeDB.GetAllReturns(data)

	serviceInstance := service.NewEmployeeService(fakeDB)
	actual := serviceInstance.GetAllEmployees()
	assert.Equal(t, data, actual)
}

func TestDeleteEmployees(t *testing.T) {
	fakeDB := &servicefakes.FakeDatabaseInterface{}

	fakeDBError := errors.New("fake DB Error")
	fakeServiceError := errors.New("no user has been deleted")
	// always works for 1, 3, and 42,
	// always db error for 0, otherwise no error, but also no deleting anything
	fakeDB.DeleteByIDCalls(func(id string) (int64, error) {
		if id == "1" || id == "3" || id == "42" {
			return 1, nil
		}
		if id == "0" {
			return 0, fakeDBError
		}
		return 0, nil
	})

	var tests = []struct {
		ids            []string
		expectedResult []string
		expectedErr    error
	}{
		{[]string{"1"}, []string{"1"}, nil},
		{[]string{"42", "3", "1"}, []string{"42", "3", "1"}, nil},
		{[]string{"0"}, nil, fakeDBError},
		{[]string{"1", "0"}, []string{"1"}, fakeDBError},
		{[]string{"1", "2"}, []string{"1"}, fakeServiceError},
		{[]string{"0", "1"}, nil, fakeDBError},
		{[]string{"2", "1"}, nil, fakeServiceError},
	}

	for _, tt := range tests {
		serviceInstance := service.NewEmployeeService(fakeDB)
		actualResult, actualErr := serviceInstance.DeleteEmployees(tt.ids)

		assert.Equal(t, tt.expectedErr, actualErr)
		assert.Equal(t, tt.expectedResult, actualResult)
	}

}

func TestCreateEmployee(t *testing.T) {
	fakeDB := &servicefakes.FakeDatabaseInterface{}
	data := []model.Employee{
		{
			ID:        "1",
			FirstName: "jon",
			LastName:  "kock",
			Email:     "jon@gmail.com",
		},
	}
	testEmployee := model.Employee{}
	fakeDB.GetByIDReturns(data[0])
	fakeDB.GetByIDReturns(testEmployee)
	fakeDB.UpdateManyReturns(data)
	serviceInstance := service.NewEmployeeService(fakeDB)
	actual, _ := serviceInstance.CreateEmployees(data)
	assert.Equal(t, data, actual)
}
