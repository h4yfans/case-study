package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/h4yfans/case-study/common/db"
	"github.com/h4yfans/case-study/common/environment"
	"github.com/h4yfans/case-study/common/logging"
	_userDelivery "github.com/h4yfans/case-study/user/delivery"
	_userRepo "github.com/h4yfans/case-study/user/repository"
	_userUsecase "github.com/h4yfans/case-study/user/usecase"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
)

type Configuration struct {
	RoutePrefix    string
	Port           int
	DB             db.Config
	ContextTimeout time.Duration
	Debug          bool
}

func main() {
	// Read .env file for local development
	environment.ReadEnvFile()

	config := Configuration{
		Port:           environment.Port(),
		DB:             environment.Database(),
		ContextTimeout: environment.ContextTimeout(),
		Debug:          environment.Debug(),
	}

	// Router
	rootRouter := mux.NewRouter()

	// Initialize logging
	logging.Initialize()
	defer logging.Close()

	// Configure Database
	boil.DebugMode = environment.BoilDebug()
	DB := db.Connect(config.DB)
	db.Migrate(DB, config.DB)
	defer db.Close(DB)

	headersOk := handlers.AllowedHeaders([]string{"content-type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	// Initialize Repositories
	// -- User --
	userRepo := _userRepo.NewUserRepository(DB)

	// Initialize Usecase
	// -- User --
	userUsecase := _userUsecase.NewUserUsecase(userRepo)

	// Initialize Handler
	_userDelivery.NewUserHandler(userUsecase, rootRouter)

	// Serve
	http.Handle("/", rootRouter)
	zap.S().Infof("Starting listening %v", config.Port)
	zap.S().Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.Port), handlers.CORS(originsOk, headersOk, methodsOk)(rootRouter)))
}
