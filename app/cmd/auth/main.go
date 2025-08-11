package main

import (
	"app/config"
	"app/internal/handlers"
	"app/internal/repository"
	"app/internal/router"
	"app/internal/services"
)

func main() {
	appConfig, _ := config.LoadConfig()
	lg := config.InitLogger()
	// _ := config.InitDB(appConfig)
	rt := config.InitGlobalLimitRate()

	// repositories
	var authRepo repository.AuthRepository
	switch appConfig.Database.Driver {
	case "postgres":
		// authRepo = repository.NewAuthRepositoryPosrtgresql(db)
	case "mysql":
		// authRepo = repository.NewAuthRepositoryMysql(db)
	case "local":
		authRepo = repository.NewAuthRepositoryLocal()
	case "grpc":
		authRepo = repository.NewAuthRepositoryGRPCClient(appConfig.Server.GrpcAddr)

	}

	// services
	authService := services.NewAuthService(authRepo)
	// handlers
	authHandlers := handlers.NewAuthHandler(authService, appConfig)

	// router
	router.InitAuthRouter(appConfig, authHandlers, rt, lg)
}
