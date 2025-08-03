package main
import ( 
    "app/config" 
    "app/internal/services" 
    "app/internal/repository"
    "app/internal/handlers"   
    // "app/internal/router"
    "app/internal/gRPCServer"
)


func main() {
    
    config.InitLogger()
    // db, config := config.InitDB()
	_, config := config.InitDB()
    // rt := config.InitGlobalLimitRate()

    // repositories
    var umsRepo repository.UmsRepository
    switch config.Driver {
    case "postgres":
        // umsRepo = repository.NewUmsRepositoryPosrtgresql(db)
    case "mysql":
        // umsRepo = repository.NewUmsRepositoryMysql(db) 
	case "local":
		umsRepo = repository.NewUmsRepositoryLocal()
	
    }

    // services
    umsService := services.NewUmsService(umsRepo)
    
    grpcSrv := gRPCServer.NewUmsGRPCServer(config.GrpcAddr )
    // handlers 
    handlers.NewGrpcUmsHandler(grpcSrv.GrpcServer,umsService)

	grpcSrv.Run()


    // router
    // router.InitUmsRouter(umsHandlers, rt, lg)
}
