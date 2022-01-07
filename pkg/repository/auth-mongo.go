package repository
//// package for database authentication functions!
import (
	structs "Golangcrud/Structs"
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authmongo struct {
	db *mongo.Client
}

func NewAuthMongo(db *mongo.Client) *Authmongo {
	return &Authmongo{db: db}
}
func (r *Authmongo) CreateUser(user structs.User) (string, error) {
	// var id int;
	db := r.db.Database("Todo").Collection("Users")
	result, err := db.InsertOne(context.TODO(), user)
	if err != nil {
		logrus.Fatalf("Error while login %s", err.Error())
	}
	hexid := result.InsertedID.(primitive.ObjectID).Hex()

	fmt.Printf("Id %v", result)
	return hexid, nil
}

func (r *Authmongo) Getuser(username, password string) (structs.User, error) {
	db := r.db.Database("Todo").Collection("Users")
	var user structs.User
	err := db.FindOne(context.TODO(),bson.D{ 
		{"username",username},{"password",password},
	}).Decode(&user);

	if err != nil{
		return user,err
	}
	return user,nil;
}
