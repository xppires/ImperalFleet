package main

import (
	"app/config"
	"app/internal/handlers"
	"app/internal/repository"
	"app/internal/services"
	"fmt"

	// "app/internal/router"
	"app/internal/gRPCServer"
)

func main() {

	appConfig, _ := config.LoadConfig()

	config.InitLogger()
	// db, config := config.InitDB()
	// _ := config.InitDB(appConfig)
	// rt := config.InitGlobalLimitRate()

	// repositories
	var umsRepo repository.UmsRepository
	switch appConfig.Database.Driver {
	case "postgres":
		// umsRepo = repository.NewUmsRepositoryPosrtgresql(db)
	case "mysql":
		// umsRepo = repository.NewUmsRepositoryMysql(db)
	case "local":
		umsRepo = repository.NewUmsRepositoryLocal()

	}

	// services
	umsService := services.NewUmsService(umsRepo)

	grpcSrv := gRPCServer.NewUmsGRPCServer(appConfig.Server.GrpcAddr)
	// handlers
	handlers.NewGrpcUmsHandler(grpcSrv.GrpcServer, umsService)

	err := grpcSrv.Run()
	if err != nil {
		fmt.Printf("grpcSrv.Run() err: %v\n", err)
	}

	// router
	// router.InitUmsRouter(umsHandlers, rt, lg)
}
