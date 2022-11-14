package main

import (
	"os"

	todoapi "github.com/AbdulahadAbduqahhorov/gin/todo-api"
	"github.com/AbdulahadAbduqahhorov/gin/todo-api/api/handlers"
	"github.com/AbdulahadAbduqahhorov/gin/todo-api/api/repository"
	"github.com/AbdulahadAbduqahhorov/gin/todo-api/api/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	err := initConfig()
	if err != nil {
		logrus.Fatalf("error initializing config: %v", err.Error())
	}
	err = godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	db, err := repository.NewPostgres(repository.Config{
		Host:     viper.GetString("postgres.host"),
		Port:     viper.GetString("postgres.port"),
		UserName: viper.GetString("postgres.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("postgres.dbname"),
		SSLMode:  viper.GetString("postgres.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("error connecting to postgres: %v", err.Error())
	}
	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handlers.NewHandler(service)
	srv := new(todoapi.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitializeHandler()); err != nil {
		logrus.Fatalf(err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	return viper.ReadInConfig()
}
