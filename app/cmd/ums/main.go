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
    var umsRepo repository.UmsRepository
    switch dbConfig.Driver {
    case "postgres":
        // umsRepo = repository.NewUmsRepositoryPosrtgresql(db)
    case "mysql":
        // umsRepo = repository.NewUmsRepositoryMysql(db) 
	case "local":
		umsRepo = repository.NewUmsRepositoryLocal()
	
    }

    // services
    umsService := services.NewUmsService(umsRepo)
    // handlers 
    umsHandlers := handlers.NewUmsHandler(umsService)

    // router
    router.InitUmsRouter(umsHandlers, rt, lg)
}
