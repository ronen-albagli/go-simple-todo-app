package config

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	adapter "todo/adapter"
)

type Config struct {
	mongo *mongo.Client
}

func (c *Config) GetMongoClient() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(900)*time.Millisecond)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	c.mongo = client
	if err != nil {
		println("NO CLIENT")
		return nil, err
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Mongo is Connected")
	return c.mongo, nil
}

func (c *Config) GetMongoInMemory() *adapter.TodoMongoInMemoryGateway {
	mongoMemory := &adapter.TodoMongoInMemoryGateway{}

	return mongoMemory
}

func (c *Config) GetMongoTodoGateway() *mongo.Collection {
	client, err := c.GetMongoClient()

	if err != nil {
		return nil
	}

	todoCollection := client.Database("Mongo").Collection("todo")

	if err != nil {
		return nil
	}

	return todoCollection
}
