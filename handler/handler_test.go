package handler_test

import (
	"encoding/json"
	"errors"
	"example-project/handler"
	"example-project/handler/handlerfakes"
	"example-project/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUpdateEmployeeHandler(t *testing.T) {

	validBody := "{\n  \"employees\": [\n    {\n      \"id\": \"1\",\n      \"first_name\": \"John\",\n      \"last_name\": \"Kenn\",\n      \"email\": \"john@gmail.com\"\n    },\n    {\n      \"id\": \"2\",\n      \"first_name\": \"Maria\",\n      \"last_name\": \"gonjaless\",\n      \"email\": \"maria@gmail.com\"\n    },\n    {\n      \"id\": \"3\",\n      \"first_name\": \"Lora\",\n      \"last_name\": \"kai\",\n      \"email\": \"lora@gmail.com\"\n    }\n  ]\n}"

	fakeServiceError := errors.New("fake service error")

	var tests = []struct {
		body             string
		serviceResponse  []string
		serviceErr       error
		expectedStatus   int
		expectedResponse string
	}{
		{"invalid body", nil, nil, http.StatusBadRequest, "{\"errorMessage\":\"invalid payload\"}"},
		{validBody, []string{"1", "2", "3"}, nil, http.StatusOK, "{\"updatedIDs\":[\"1\",\"2\",\"3\"]}"},
		{validBody, []string{"1", "2", "3"}, fakeServiceError, http.StatusInternalServerError, "{\"updatedIDs\":[\"1\",\"2\",\"3\"]}"},
	}

	for _, tt := range tests {
		fakeService := &handlerfakes.FakeServiceInterface{}
		fakeService.UpdateEmployeesReturns(tt.serviceResponse, tt.serviceErr)

		responseRecorder := httptest.NewRecorder()

		fakeContext, _ := gin.CreateTestContext(responseRecorder)

		fakeContext.Request = httptest.NewRequest("POST", "http://localhost:9090/employee/update", strings.NewReader(tt.body))

		handlerInstance := handler.NewHandler(fakeService)
		handlerInstance.UpdateEmployeeHandler(fakeContext)

		assert.Equal(t, tt.expectedStatus, responseRecorder.Code)
		assert.Equal(t, tt.expectedResponse, responseRecorder.Body.String())
	}
}

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

func TestCreateEmployeeHandler(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeService.CreateEmployeesReturns("myFantasyString", nil)

	payload := model.Payload{
		Employees: []model.Employee{
			{
				ID:        "1",
				FirstName: "John",
				LastName:  "Kenn",
				Email:     "john@gmail.com",
			},
		},
	}
	jsonPayload, _ := json.Marshal(payload)

	fakeContext, _ := gin.CreateTestContext(responseRecorder)
	fakeContext.Request = httptest.NewRequest("POST", "http://localhost:9090/employee/create", strings.NewReader(string(jsonPayload)))

	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.CreateEmployeeHandler(fakeContext)

	assert.Equal(t, 200, responseRecorder.Code)

}
