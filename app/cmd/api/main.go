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
    db := config.InitMysql()

    // repositories
    spacecraftRepo := repository.NewSpacecraftRepository(db)

    // services
    spacecraftService := services.NewSpacecraftService(spacecraftRepo)
   // handlers 
    spacecraftHandlers := handlers.NewSpacecrafHandlers(spacecraftService)

    // router
    router.InitRouter(spacecraftHandlers, lg)
}
