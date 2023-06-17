package adapter

import (
	"fmt"
	"todo/gateway"
	types "todo/types"
	// "go.mongodb.org/mongo-driver/mongo"
)

type TodoMongoInMemoryGateway struct {
	Collection interface{}
}

func (m *TodoMongoInMemoryGateway) Store(todo types.TodoEvent) error {
	fmt.Println("Saved in memory")

	gateway.Save()

	return nil
}
