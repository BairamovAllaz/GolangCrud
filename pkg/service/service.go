package service

import (
	structs "Golangcrud/Structs"
	"Golangcrud/pkg/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	CreateUser(user structs.User)(string,error)
	GenerateToken(username,password string)(string,error)
	Parsetoken(token string)(int,error)
	ForgotMypassword(myuser string,username structs.Fpasswordstruct)(string,error)
	Checkdatabaseusertoken(token string)(structs.User,error)
	ChangePassword(user structs.User,password structs.Newpassword)(string,error)
}
type Todolist interface {
	Create(UserId int,list structs.Todolist)(string,error)
	GetAll(UserId int)([]structs.Todolist,error)
	GetListById(id string)(structs.Todolist,error);
	UpdateList(input structs.UpdateListItem,id string)(*mongo.UpdateResult,error)
	DeleteList(id string)(*mongo.DeleteResult,error)
}
type Todoitem interface {
}
type Service struct {
	Authorization
	Todolist
	Todoitem
}

///constructor

func NewServise(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repos.Authorization),
		Todolist: NewTodolistService(repos.Todolist),
	}
}
