package usecase

import (
	"testing"
	"todo/adapter"
	"todo/config"
	"todo/types"
)

func TestAddTodo(t *testing.T) {
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

	todoId, err := usecase.Do(*input)

	if err != nil {
		t.Error("Use case failed, Error: ", err)

		return
	}

	if todoId != "" {
		t.Log("Test passed")
	}
}
