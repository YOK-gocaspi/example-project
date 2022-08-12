package datasource

import (
	"context"
	"example-project/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . MongoDBInterface
type MongoDBInterface interface {
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
	DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
}

type Client struct {
	Employee MongoDBInterface
}

func NewDbClient(d model.DbConfig) Client {
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI(d.URL))
	db := client.Database(d.Database)
	return Client{
		Employee: db.Collection("employee"),
	}
}

func (c Client) UpdateMany(docs []interface{}) interface{} {
	results, err := c.Employee.InsertMany(context.TODO(), docs)
	if err != nil {
		log.Println("database error")
	}
	return results.InsertedIDs
}

func (c Client) DeleteByID(id string) (int64, error) {
	filter := bson.D{{"id", id}}

	results, err := c.Employee.DeleteMany(context.TODO(), filter)
	if err != nil {
		return results.DeletedCount, err
	}
	return results.DeletedCount, nil
}

func (c Client) GetByID(id string) model.Employee {
	filter := bson.M{"id": id}
	courser := c.Employee.FindOne(context.TODO(), filter)
	var employee model.Employee
	err := courser.Decode(&employee)
	if err != nil {
		log.Println("error during data marshalling")
	}
	return employee
}

func (c Client) GetAll() []model.Employee {
	filter := bson.M{}
	cursor, err := c.Employee.Find(context.TODO(), filter)

	var employees []model.Employee
	err2 := cursor.All(context.TODO(), &employees)
	if err != nil {
		log.Println("error during data marshalling")
	}
	if err2 != nil {
		log.Println("error:", err2)
	}

	return employees
}
