package ports

import (
	types "todo/types"
)

type TodoGateway interface {
	Store(todo types.TodoEvent) error
}
