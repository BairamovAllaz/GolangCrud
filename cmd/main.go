package main

import (
	"Golangcrud/Server"
	"Golangcrud/pkg/handler"
	"Golangcrud/pkg/repository"
	"Golangcrud/pkg/service"
	"os"

	// "path"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//@title TODO APP API
//@version 1.0
//@desctription My test api server
//@host localhost:8000
//@BasePath /
//@securityDefinitions.apikey ApiKeyAuth
//@in header
//@name Authorization 


func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter)) //set logrus formatter~
	if err := initConfig(); err != nil {
		logrus.Fatalf("error while initilazion config: %s", err.Error())
	}

	if err := godotenv.Load(filepath.Join(".", ".env")); err != nil {
		logrus.Fatalf("Error while loaded dotenv file %s", err.Error())
	}

	db, err := repository.NewMongoDb(&repository.Config{ //connect to database mongoDB
		Uri: os.Getenv("DB_URI"),
	})

	if err != nil {
		logrus.Fatalf("Error while connect to databse %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewServise(repos)
	handlers := handler.NewHandler(services)


	
	srv := new(Server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.Initroutes()); err != nil {
		logrus.Fatalf("Error while runing http sever %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
