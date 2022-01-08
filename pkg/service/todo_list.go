package service

import (
	structs "Golangcrud/Structs"
	"Golangcrud/pkg/repository"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
)

type TodolistService struct {
	repo repository.Todolist
}

func NewTodolistService(repo repository.Todolist) *TodolistService {
	return &TodolistService{repo: repo}
}

func (s *TodolistService) Create(UserId int, list structs.Todolist) (string, error) {
	list.Id = UserId
	return s.repo.Create(UserId, list)
}
func (s *TodolistService) GetAll(UserId int) ([]structs.Todolist, error) {
	return s.repo.GetAll(UserId)
}

func (s *TodolistService) GetListById(id string) (structs.Todolist, error) {
	newid, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalf("Error %s", err.Error())
	}
	return s.repo.GetListById(newid)
}

func (s *TodolistService) UpdateList(input structs.Todolist, id string) (*mongo.UpdateResult, error) {
	newid, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalf("Error %s", err.Error())
	}
	return s.repo.UpdateList(input, newid)
}

func (s *TodolistService) DeleteList(id string) (*mongo.DeleteResult, error) {
	newid, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalf("Error %s", err.Error())
	}
	return s.repo.DeleteList(newid)
}
