package main
import ( 
    "app/config" 
    "app/internal/services" 
    "app/internal/repository"
    "app/internal/handlers"   
    "app/internal/router"
)


func main() {
    
    lg := config.InitLogger()
    db, dbConfig := config.InitDB()

    // repositories
    var spacecraftRepo repository.SpacecraftRepository
    switch dbConfig.Driver {
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
    router.InitRouter(spacecraftHandlers, lg)
}
