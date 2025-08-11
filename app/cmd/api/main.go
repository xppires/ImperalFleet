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
	db := config.InitDB(appConfig)
	rt := config.InitGlobalLimitRate()

	// repositories
	var spacecraftRepo repository.SpacecraftRepository
	switch appConfig.Database.Driver {
	case "postgres":
		spacecraftRepo = repository.NewSpacecraftRepositoryPosrtgresql(db)
	case "mysql":
		spacecraftRepo = repository.NewSpacecraftRepositoryMysql(db)
	}

	// services
	spacecraftService := services.NewSpacecraftService(spacecraftRepo)
	// handlers
	spacecraftHandlers := handlers.NewSpacecrafHandlers(spacecraftService)

	// router
	router.InitRouter(appConfig, spacecraftHandlers, rt, lg)
}
