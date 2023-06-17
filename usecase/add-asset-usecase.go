package usecase

import (
	"todo/ports"
	"todo/types"
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

func (addUseCase AddTodo) Do(todo types.TodoEvent) error {
	c := addUseCase.Config

	err := c.TodoGateway.Store(todo)

	return err
}
