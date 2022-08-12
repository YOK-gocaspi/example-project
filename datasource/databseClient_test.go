package datasource_test

import (
	"context"
	"example-project/datasource"
	"example-project/datasource/datasourcefakes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
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

func TestGetAllEmployees(t *testing.T) {
	fakeDB := &datasourcefakes.FakeMongoDBInterface{}
	// fakeDB.InsertManyReturns(&mongo.InsertManyResult{InsertedIDs: []interface{}{"1"}}, nil)
	cursor, _ := fakeDB.Find(context.TODO(), bson.M{})
	fakeDB.FindReturns(cursor, nil)
	dbClient := datasource.Client{
		Employee: fakeDB,
	}
	result := dbClient.GetAll()
	fmt.Println("Get All:", result)
	/*
		// DB erstellen -> Daten schreiben -> Input == Output

		dbClient := datasource.Client{
			Employee: fakeDB,
		}
		actual := dbClient.GetAll()
	*/
}
