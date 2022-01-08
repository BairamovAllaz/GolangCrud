package repository

import (
	structs "Golangcrud/Structs"

	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	CreateUser(user structs.User) (string, error)
	Getuser(username, password string) (structs.User, error)
}
type Todolist interface {
	Create(UserId int, list structs.Todolist) (string, error)
	GetAll(UserId int) ([]structs.Todolist, error)
	GetListById(id int) (structs.Todolist, error)
	UpdateList(input structs.UpdateListItem, id int) (*mongo.UpdateResult, error)
	DeleteList(id int)(*mongo.DeleteResult,error)
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
		Todolist:      NewTodoMongoRepo(db),
	}
}
