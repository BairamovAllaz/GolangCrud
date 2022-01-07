package service

import (
	structs "Golangcrud/Structs"
	"Golangcrud/pkg/repository"
)

type Authorization interface {
	CreateUser(user structs.User)(string,error)
	GenerateToken(username,password string)(string,error)
	Parsetoken(token string)(int,error)
}
type Todolist interface {
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
	}
}
