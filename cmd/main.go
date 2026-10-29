package main

import (
	"log"

	todo "github.com/Iogamy/TODO-Project"
	"github.com/Iogamy/TODO-Project/pkg/handler"
	"github.com/Iogamy/TODO-Project/pkg/repository"
	"github.com/Iogamy/TODO-Project/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepoository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error)
	}
}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
