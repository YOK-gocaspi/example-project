package datasource_test

import (
	"errors"
	"example-project/datasource"
	"example-project/datasource/datasourcefakes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestUpdateMany(t *testing.T) {
	fakeDb := &datasourcefakes.FakeMongoDBInterface{}
	fakeDb.InsertManyReturns(&mongo.InsertManyResult{InsertedIDs: []interface{}{"1"}}, nil)
	dbClient := datasource.Client{
		Employee: fakeDb,
	}
	actual := dbClient.UpdateMany([]interface{}{})
	fmt.Println(actual)
	assert.NotEmpty(t, actual)
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
