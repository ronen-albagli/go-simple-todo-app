package adapter

import (
	"context"
	"errors"
	"fmt"

	types "todo/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoMongoGateway struct {
	mongo.Collection
}

func (m *TodoMongoGateway) Store(todo types.TodoEvent) error {
	ctx := context.Background()

	data, err := m.InsertOne(
		ctx,
		bson.M{
			"name":        todo.Name,
			"description": todo.Description,
			"status":      todo.Status,
		})

	if err != nil {
		return errors.New("failed to insert to ledger")
	}

	fmt.Println(data)

	return nil
}
