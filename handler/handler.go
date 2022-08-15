package handler

import (
	"example-project/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . ServiceInterface
type ServiceInterface interface {
	CreateEmployees(employees []model.Employee) interface{}
	UpdateEmployees(employees []model.Employee) ([]string, error)
	GetEmployeeById(id string) model.Employee
	GetAllEmployees() []model.Employee
	DeleteEmployees(ids []string) (interface{}, error)
}

type Handler struct {
	ServiceInterface ServiceInterface
}

func NewHandler(serviceInterface ServiceInterface) Handler {
	return Handler{
		ServiceInterface: serviceInterface,
	}
}

func (handler Handler) CreateEmployeeHandler(c *gin.Context) {
	var payLoad model.Payload
	err := c.BindJSON(&payLoad)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorMessage": "invalid payload",
		})
		return
	}

	response := handler.ServiceInterface.CreateEmployees(payLoad.Employees)
	c.JSON(200, response)
}

func (handler Handler) UpdateEmployeeHandler(c *gin.Context) {
	var payLoad model.Payload
	err := c.BindJSON(&payLoad)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorMessage": "invalid payload",
		})
		return
	}

	updatedIDs, err := handler.ServiceInterface.UpdateEmployees(payLoad.Employees)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"updatedIDs": updatedIDs,
		})
		return
	}
	c.JSON(200, gin.H{
		"updatedIDs": updatedIDs,
	})
}

func (handler Handler) GetEmployeeHandler(c *gin.Context) {
	pathParam, ok := c.Params.Get("id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorMessage": "id is not given",
		})
		return
	}

	response := handler.ServiceInterface.GetEmployeeById(pathParam)
	c.JSON(http.StatusOK, response)
}

func (handler Handler) GetAllEmployeesHandler(c *gin.Context) {
	response := handler.ServiceInterface.GetAllEmployees()
	c.JSON(http.StatusOK, response)
}

func (handler Handler) DeleteEmployeeHandler(c *gin.Context) {
	ids := c.QueryArray("id")
	response, err := handler.ServiceInterface.DeleteEmployees(ids)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response)

	} else {
		c.JSON(http.StatusOK, response)
	}
}
