package handler_test

import (
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
