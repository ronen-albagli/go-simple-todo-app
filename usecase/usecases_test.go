package usecase

import (
	"testing"
	"todo/adapter"
	"todo/config"
	"todo/types"

	"go.mongodb.org/mongo-driver/mongo"
)

type AddLedgerEvent struct {
	ledgerCollection *mongo.Collection
}

func TestAddAssets(t *testing.T) {
	var input = &types.TodoEvent{}

	config := new(config.Config)

	configuration := Config{
		TodoGateway: &adapter.TodoMongoInMemoryGateway{
			Collection: *config.GetMongoInMemory(),
		},
	}

	usecase := &AddTodo{
		Config: configuration,
	}

	input.Name = "test todo"
	input.Description = "test description"
	input.Status = false

	err := usecase.Do(*input)

	if err != nil {
		t.Error("Use case failed, Error: ", err)

		return
	}

	t.Log("Test passed")

}

func TestHelloEmpty(t *testing.T) {

}
