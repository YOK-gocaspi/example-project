package datasource_test

import (
	"errors"
	"example-project/datasource"
	"example-project/datasource/datasourcefakes"
	"example-project/model"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestCreateMany(t *testing.T) {
	fakeDb := &datasourcefakes.FakeMongoDBInterface{}
	fakeDb.InsertManyReturns(&mongo.InsertManyResult{InsertedIDs: []interface{}{"1"}}, nil)
	dbClient := datasource.Client{
		Employee: fakeDb,
	}
	actual := dbClient.UpdateMany([]interface{}{})
	fmt.Println(actual)
	assert.NotEmpty(t, actual)
}

func TestUpdateEmployee(t *testing.T) {
	testEmployee := model.Employee{
		ID:        "1",
		FirstName: "Frank",
		LastName:  "F",
		Email:     "F@test.com",
	}

	fakeResultSucess := mongo.UpdateResult{
		MatchedCount:  1,
		ModifiedCount: 1,
		UpsertedCount: 0,
		UpsertedID:    nil,
	}
	fakeResultNotInDatabase := mongo.UpdateResult{
		MatchedCount:  0,
		ModifiedCount: 0,
		UpsertedCount: 0,
		UpsertedID:    nil,
	}
	fakeDBError := errors.New("database error")
	fakeErrNotInDB := errors.New("no user with id " + "1" + " in database")

	var tests = []struct {
		dbresult    mongo.UpdateResult
		dbErr       error
		expectedErr error
	}{
		{fakeResultSucess, nil, nil},
		{fakeResultNotInDatabase, nil, fakeErrNotInDB},
		{fakeResultNotInDatabase, fakeDBError, fakeDBError},
	}

	for _, tt := range tests {
		fakeDb := &datasourcefakes.FakeMongoDBInterface{}
		fakeDb.UpdateOneReturns(&tt.dbresult, tt.dbErr)
		dbClient := datasource.Client{
			Employee: fakeDb,
		}
		actualErr := dbClient.UpdateEmployee(testEmployee)
		assert.Equal(t, tt.expectedErr, actualErr)
	}
}

func TestDeleteByID(t *testing.T) {
	fakeDBError := errors.New("fake DB Error")

	var tests = []struct {
		dbDeletedCount int64
		dbErr          error
		expectedResult int64
		expectedErr    error
	}{{1, nil, 1, nil},
		{0, fakeDBError, 0, fakeDBError},
	}

	for _, tt := range tests {
		fakeDb := &datasourcefakes.FakeMongoDBInterface{}
		fakeDb.DeleteManyReturns(&mongo.DeleteResult{DeletedCount: tt.dbDeletedCount}, tt.dbErr)
		dbClient := datasource.Client{
			Employee: fakeDb,
		}
		actualResult, actualErr := dbClient.DeleteByID("1")
		assert.Equal(t, tt.expectedErr, actualErr)
		assert.Equal(t, tt.expectedResult, actualResult)
	}

}
