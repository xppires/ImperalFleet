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
    // db, dbConfig := config.InitDB()
	_, dbConfig := config.InitDB()
    rt := config.InitGlobalLimitRate()

    // repositories
    var authRepo repository.AuthRepository
    switch dbConfig.Driver {
    case "postgres":
        // authRepo = repository.NewAuthRepositoryPosrtgresql(db)
    case "mysql":
        // authRepo = repository.NewAuthRepositoryMysql(db) 
	case "local":
		authRepo = repository.NewAuthRepositoryLocal()
	case "grpc":
		authRepo = repository.NewAuthRepositoryGRPCClient(dbConfig.GrpcAddr)
	
    }

    // services
    authService := services.NewAuthService(authRepo)
    // handlers 
    authHandlers := handlers.NewAuthHandler(authService)

    // router
    router.InitAuthRouter(authHandlers, rt, lg)
}
