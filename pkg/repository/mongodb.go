package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct{ 
	Uri string 
}

func NewMongoDb(cfg *Config) (*mongo.Client,error){ 
	client,err := mongo.Connect(context.TODO(),options.Client().ApplyURI(cfg.Uri));
	if err != nil{ 
		return nil,err;
	}
	return client, nil;
}
