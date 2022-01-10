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
	db := r.db.Database("Todo").Collection("Users")
	result, err := db.InsertOne(context.TODO(), user)
	if err != nil {
		logrus.Fatalf("Error while login %s", err.Error())
	}
	hexid := result.InsertedID.(primitive.ObjectID).Hex()
	fmt.Printf("Id %v", result)
	return hexid, nil
}

func(r *Authmongo)ForgotMypassword(token string,myuser string,username structs.Fpasswordstruct)(string,error){ 
	db := r.db.Database("Todo").Collection("Users")

	filter := bson.D{{"username",username.Username}}; 
	updateelement := bson.D{{"$set",bson.D{{"token",token}}}}; 

	_,err := db.UpdateOne(context.TODO(),filter,updateelement); 
	
	if err != nil { 
		return "",err;
	}
	return token,nil;
}



func (r *Authmongo)Checkdatabaseusertoken(token string)(structs.User,error){ 
	db := r.db.Database("Todo").Collection("Users");
	filter := bson.D{{"token",token}}; 
	var user structs.User;
	err := db.FindOne(context.TODO(),filter).Decode(&user); 
	if err != nil {
		return structs.User{},err;
	}
	return user,nil;
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

func (r *Authmongo)ChangePassword(username string,newpassword string)(string,error){ 
	db := r.db.Database("Todo").Collection("Users");
	filter := bson.D{{"username",username}}; 
	updateelement := bson.D{{"$set",bson.D{{"password",newpassword}}}}; 
	_,err := db.UpdateOne(context.TODO(),filter,updateelement);
	if err != nil { 
		return "",err;
	}
	cleartoken := bson.D{{"$set",bson.D{{"token",""}}}}
	_,ok := db.UpdateOne(context.TODO(),filter,cleartoken);
	if ok != nil{ 
		return "",nil;
	}
	return "Password update succes and clear token success",nil;
}
