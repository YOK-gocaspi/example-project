package datasource_test

import (
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
