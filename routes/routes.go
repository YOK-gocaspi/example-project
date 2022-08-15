package routes

import (
	"github.com/gin-gonic/gin"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . HandlerInterface
type HandlerInterface interface {
	CreateEmployeeHandler(c *gin.Context)
	UpdateEmployeeHandler(c *gin.Context)
	GetEmployeeHandler(c *gin.Context)
	GetAllEmployeesHandler(c *gin.Context)
	DeleteEmployeeHandler(c *gin.Context)
}

var Handler HandlerInterface

func CreateRoutes(group *gin.RouterGroup) {
	route := group.Group("/employee")
	route.GET("/:id/get", Handler.GetEmployeeHandler)
	route.GET("/get", Handler.GetAllEmployeesHandler)
	route.POST("/create", Handler.CreateEmployeeHandler)
	route.POST("/update", Handler.UpdateEmployeeHandler)
	route.DELETE("/delete", Handler.DeleteEmployeeHandler)

}
