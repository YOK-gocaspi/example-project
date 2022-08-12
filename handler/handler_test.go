package handler_test

import (
	"errors"
	"example-project/handler"
	"example-project/handler/handlerfakes"
	"example-project/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEmployeeHandler_Return_valid_status_code(t *testing.T) {
	responseRecoder := httptest.NewRecorder()

	fakeContest, _ := gin.CreateTestContext(responseRecoder)
	fakeContest.Params = append(fakeContest.Params, gin.Param{Key: "id", Value: "1"})

	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeService.GetEmployeeByIdReturns(model.Employee{
		ID:        "1",
		FirstName: "Joe",
	})

	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.GetEmployeeHandler(fakeContest)

	assert.Equal(t, http.StatusOK, responseRecoder.Code)

}

func TestGetAllEmployees(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeService.GetAllEmployeesCalls(func() []model.Employee {
		return []model.Employee{
			{},
		}
	})

	fakeContext, _ := gin.CreateTestContext(responseRecorder)
	fakeContext.Request = httptest.NewRequest("GET", "http://localhost:9090/employee/get", nil)

	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.GetAllEmployeesHandler(fakeContext)

	assert.Equal(t, 200, responseRecorder.Code)
}

func TestDeleteEmployeeHandler(t *testing.T) {

	fakeServiceError := errors.New("fake service error")

	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeService.DeleteEmployeesCalls(func(ids []string) (interface{}, error) {
		var deletedIDs []string
		for _, id := range ids {
			if id == "0" {
				return deletedIDs, fakeServiceError
			}
			deletedIDs = append(deletedIDs, id)
		}
		return deletedIDs, nil
	})

	var tests = []struct {
		ids              []string
		expectedStatus   int
		expectedResponse string
	}{
		{[]string{"0"}, http.StatusInternalServerError, "null"},
		{[]string{"1"}, http.StatusOK, "[\"1\"]"},
		{[]string{"1", "0"}, http.StatusInternalServerError, "[\"1\"]"},
		{[]string{"1", "2", "3"}, http.StatusOK, "[\"1\",\"2\",\"3\"]"},
	}

	for _, tt := range tests {
		responseRecorder := httptest.NewRecorder()

		fakeContext, _ := gin.CreateTestContext(responseRecorder)
		var query string
		for _, id := range tt.ids {
			if query != "" {
				query = query + "&"
			}
			query = query + "id=" + id
		}
		fakeContext.Request = httptest.NewRequest("DELETE", "http://localhost:9090/employee/delete?"+query, nil)

		handlerInstance := handler.NewHandler(fakeService)
		handlerInstance.DeleteEmployeeHandler(fakeContext)

		assert.Equal(t, tt.expectedStatus, responseRecorder.Code)
		assert.Equal(t, tt.expectedResponse, responseRecorder.Body.String())
	}

}
