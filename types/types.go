package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type INewTodoInput interface {
	TodoInput
}

type TodoInput struct {
	Name        string
	Description string
	Status      bool
}

type TodoEvent struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
	Status      bool               `bson:"status,omitempty"`
}
