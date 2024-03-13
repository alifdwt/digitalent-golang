package main

import (
	"log"

	"github.com/alifdwt/digitalent-golang/handler"
	"github.com/alifdwt/digitalent-golang/mapper"
	"github.com/alifdwt/digitalent-golang/pkg/database/migration"
	"github.com/alifdwt/digitalent-golang/pkg/database/postgres"
	"github.com/alifdwt/digitalent-golang/pkg/dotenv"
	"github.com/alifdwt/digitalent-golang/repository"
	"github.com/alifdwt/digitalent-golang/service"
)

// @title Digitalent-Golang
// @version 1.0
// @description This is an API for Digitalent-Golang Assignment 2
// @termsOfService http://swagger.io/terms/

// @contact.name Alif Dewantara
// @contact.url http://github.com/alifdwt
// @contact.email aputradewantara@gmail.com

// @host localhost:8080
func main() {
	config, err := dotenv.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	db, err := postgres.NewClient(config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_NAME)
	if err != nil {
		log.Fatal("Error while connecting to database: ", err)
	}

	err = migration.RunMigration(db)
	if err != nil {
		log.Fatal("Error while migrating to database: ", err)
	}

	repository := repository.NewRepository(db)

	mapper := mapper.NewMapper()

	service := service.NewService(service.Deps{
		Repositories: repository,
		Mapper:       *mapper,
	})

	myHandler, err := handler.NewHandler(*service)
	if err != nil {
		log.Fatal("Error while creating server: ", err)
	}

	err = myHandler.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
