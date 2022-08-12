// Code generated by counterfeiter. DO NOT EDIT.
package routesfakes

import (
	"example-project/routes"
	"sync"

	"github.com/gin-gonic/gin"
)

type FakeHandlerInterface struct {
	CreateEmployeeHandlerStub        func(*gin.Context)
	createEmployeeHandlerMutex       sync.RWMutex
	createEmployeeHandlerArgsForCall []struct {
		arg1 *gin.Context
	}
	GetAllEmployeesHandlerStub        func(*gin.Context)
	getAllEmployeesHandlerMutex       sync.RWMutex
	getAllEmployeesHandlerArgsForCall []struct {
		arg1 *gin.Context
	}
	GetEmployeeHandlerStub        func(*gin.Context)
	getEmployeeHandlerMutex       sync.RWMutex
	getEmployeeHandlerArgsForCall []struct {
		arg1 *gin.Context
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHandlerInterface) CreateEmployeeHandler(arg1 *gin.Context) {
	fake.createEmployeeHandlerMutex.Lock()
	fake.createEmployeeHandlerArgsForCall = append(fake.createEmployeeHandlerArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.CreateEmployeeHandlerStub
	fake.recordInvocation("CreateEmployeeHandler", []interface{}{arg1})
	fake.createEmployeeHandlerMutex.Unlock()
	if stub != nil {
		fake.CreateEmployeeHandlerStub(arg1)
	}
}

func (fake *FakeHandlerInterface) CreateEmployeeHandlerCallCount() int {
	fake.createEmployeeHandlerMutex.RLock()
	defer fake.createEmployeeHandlerMutex.RUnlock()
	return len(fake.createEmployeeHandlerArgsForCall)
}

func (fake *FakeHandlerInterface) CreateEmployeeHandlerCalls(stub func(*gin.Context)) {
	fake.createEmployeeHandlerMutex.Lock()
	defer fake.createEmployeeHandlerMutex.Unlock()
	fake.CreateEmployeeHandlerStub = stub
}

func (fake *FakeHandlerInterface) CreateEmployeeHandlerArgsForCall(i int) *gin.Context {
	fake.createEmployeeHandlerMutex.RLock()
	defer fake.createEmployeeHandlerMutex.RUnlock()
	argsForCall := fake.createEmployeeHandlerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) GetAllEmployeesHandler(arg1 *gin.Context) {
	fake.getAllEmployeesHandlerMutex.Lock()
	fake.getAllEmployeesHandlerArgsForCall = append(fake.getAllEmployeesHandlerArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.GetAllEmployeesHandlerStub
	fake.recordInvocation("GetAllEmployeesHandler", []interface{}{arg1})
	fake.getAllEmployeesHandlerMutex.Unlock()
	if stub != nil {
		fake.GetAllEmployeesHandlerStub(arg1)
	}
}

func (fake *FakeHandlerInterface) GetAllEmployeesHandlerCallCount() int {
	fake.getAllEmployeesHandlerMutex.RLock()
	defer fake.getAllEmployeesHandlerMutex.RUnlock()
	return len(fake.getAllEmployeesHandlerArgsForCall)
}

func (fake *FakeHandlerInterface) GetAllEmployeesHandlerCalls(stub func(*gin.Context)) {
	fake.getAllEmployeesHandlerMutex.Lock()
	defer fake.getAllEmployeesHandlerMutex.Unlock()
	fake.GetAllEmployeesHandlerStub = stub
}

func (fake *FakeHandlerInterface) GetAllEmployeesHandlerArgsForCall(i int) *gin.Context {
	fake.getAllEmployeesHandlerMutex.RLock()
	defer fake.getAllEmployeesHandlerMutex.RUnlock()
	argsForCall := fake.getAllEmployeesHandlerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) GetEmployeeHandler(arg1 *gin.Context) {
	fake.getEmployeeHandlerMutex.Lock()
	fake.getEmployeeHandlerArgsForCall = append(fake.getEmployeeHandlerArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.GetEmployeeHandlerStub
	fake.recordInvocation("GetEmployeeHandler", []interface{}{arg1})
	fake.getEmployeeHandlerMutex.Unlock()
	if stub != nil {
		fake.GetEmployeeHandlerStub(arg1)
	}
}

func (fake *FakeHandlerInterface) GetEmployeeHandlerCallCount() int {
	fake.getEmployeeHandlerMutex.RLock()
	defer fake.getEmployeeHandlerMutex.RUnlock()
	return len(fake.getEmployeeHandlerArgsForCall)
}

func (fake *FakeHandlerInterface) GetEmployeeHandlerCalls(stub func(*gin.Context)) {
	fake.getEmployeeHandlerMutex.Lock()
	defer fake.getEmployeeHandlerMutex.Unlock()
	fake.GetEmployeeHandlerStub = stub
}

func (fake *FakeHandlerInterface) GetEmployeeHandlerArgsForCall(i int) *gin.Context {
	fake.getEmployeeHandlerMutex.RLock()
	defer fake.getEmployeeHandlerMutex.RUnlock()
	argsForCall := fake.getEmployeeHandlerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createEmployeeHandlerMutex.RLock()
	defer fake.createEmployeeHandlerMutex.RUnlock()
	fake.getAllEmployeesHandlerMutex.RLock()
	defer fake.getAllEmployeesHandlerMutex.RUnlock()
	fake.getEmployeeHandlerMutex.RLock()
	defer fake.getEmployeeHandlerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHandlerInterface) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ routes.HandlerInterface = new(FakeHandlerInterface)