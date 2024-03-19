package main

import (
	"mygram-api/handler"
	"mygram-api/mapper"
	"mygram-api/pkg/database/migration"
	"mygram-api/pkg/database/postgres"
	"mygram-api/pkg/dotenv"
	"mygram-api/pkg/hashing"
	"mygram-api/pkg/logger"
	"mygram-api/pkg/token"
	"mygram-api/repository"
	"mygram-api/service"

	"go.uber.org/zap"
)

// @title MyGram
// @version 1.0
// @description This is MyGram API for final project at Scalable Webservice with Golang - DTS Kominfo
// @termsOfService http://swagger.io/terms/

// @contact.name Alif Dewantara
// @contact.url http://github.com/alifdwt
// @contact.email aputradewantara@gmail.com

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @license.name Apache 2.0
// @licence.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
func main() {
	log, err := logger.NewLogger()
	if err != nil {
		log.Error("Error while initializing logger", zap.Error(err))
	}

	config, err := dotenv.LoadConfig(".")
	if err != nil {
		log.Error("Error while load environtment variables", zap.Error(err))
	}

	db, err := postgres.NewClient(config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_NAME)
	if err != nil {
		log.Error("Error while connecting to database", zap.Error(err))
	}

	err = migration.RunMigration(db)
	if err != nil {
		log.Error("Error while migrating to database", zap.Error(err))
	}

	hashing := hashing.NewHashingPassword()
	repository := repository.NewRepository(db)

	tokenMaker, err := token.NewJWTMaker(config.TOKEN_SYMETRIC_KEY)
	if err != nil {
		log.Error("Error while initializing token maker", zap.Error(err))
	}

	mapper := mapper.NewMapper()

	service := service.NewService(service.Deps{
		Config:     config,
		Repository: repository,
		Hashing:    *hashing,
		TokenMaker: tokenMaker,
		Logger:     *log,
		Mapper:     *mapper,
	})

	myHandler := handler.NewHandler(service, tokenMaker)

	err = myHandler.Init().Run(config.ServerAddress)
	if err != nil {
		log.Error("Cannot start server", zap.Error(err))
	}
}
