package repository

import (
	structs "Golangcrud/Structs"
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoMongoRepo struct {
	db *mongo.Client
}

func NewTodoMongoRepo(db *mongo.Client) *TodoMongoRepo {
	return &TodoMongoRepo{db: db}
}

func (s *TodoMongoRepo) Create(UserId int, list structs.Todolist) (string, error) {
	database := s.db.Database("Todo").Collection("List")
	result, err := database.InsertOne(context.TODO(), list)
	if err != nil {
		logrus.Fatalf("Error while init list %s", err.Error())
	}
	hexid := result.InsertedID.(primitive.ObjectID).Hex()
	return hexid, nil
}

func (s *TodoMongoRepo) GetAll(UserId int) ([]structs.Todolist, error) {
	var lists []structs.Todolist
	database := s.db.Database("Todo").Collection("List")

	items, err := database.Find(context.TODO(), bson.D{})
	if err != nil {
		logrus.Fatalf("Error while get all lists: %s", err.Error())
	}

	if err := items.All(context.TODO(), &lists); err != nil {
		logrus.Fatalf("Error while get all lists: %s", err.Error())
	}
	return lists, nil
}

func (s *TodoMongoRepo) GetListById(id int) (structs.Todolist, error) {
	database := s.db.Database("Todo").Collection("List")
	var result structs.Todolist
	err := database.FindOne(context.TODO(), bson.D{
		{"id", id},
	}).Decode(&result)

	if err != nil {
		return structs.Todolist{}, err
	}
	return result, nil
}

func (s *TodoMongoRepo) UpdateList(input structs.Todolist, id int) (*mongo.UpdateResult, error) {
	database := s.db.Database("Todo").Collection("List")
	
	filter := bson.D{{"id",id}};
	updateList := bson.D{{"$set",bson.D{
		{"title",input.Title},
		{"desctription",input.Desctription},
	}}}
	update,err := database.UpdateOne(context.TODO(),filter,updateList)
	if err != nil {
		return &mongo.UpdateResult{},err;
	}
	return update,nil;
}

func(s *TodoMongoRepo)DeleteList(id int)(*mongo.DeleteResult,error) {
	database := s.db.Database("Todo").Collection("List");

	filter := bson.D{{"id",id}};

	deletelist,err := database.DeleteOne(context.TODO(),filter)

	if err != nil {
		return deletelist,err;
	}

	return deletelist,nil;

}
