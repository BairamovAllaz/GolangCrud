package repository

import (
	structs "Golangcrud/Structs"

	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	CreateUser(user structs.User) (string,error)
	Getuser(username,password string) (structs.User,error)
}
type Todolist interface {
}
type Todoitem interface {
}
type Repository struct {
	Authorization
	Todolist
	Todoitem
}

///constructor

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Authorization: NewAuthMongo(db),
	}
}
