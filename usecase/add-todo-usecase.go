package usecase

import (
	"todo/ports"
	"todo/types"

	"github.com/google/uuid"
)

type Config struct {
	TodoGateway ports.TodoGateway
}

type AddTodo struct {
	IAddTodo
	Config
}

type IAddTodo interface {
	Do(asset types.TodoEvent) (string, error)
}

func (addUseCase AddTodo) Do(todo types.TodoEvent) (string, error) {
	c := addUseCase.Config

	todo.TodoId = uuid.New().String()
	err := c.TodoGateway.Store(todo)

	return todo.TodoId, err
}
