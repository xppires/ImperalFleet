package router

import (
	"log"
	"net/http" 
    "github.com/gorilla/mux"
	"app/internal/middleware"
    "app/internal/interfaces"  
	"app/internal/handlers"
	"app/config"

 
)

func InitRouter(configApp *config.ConfigApp,spacecraftHandler *handlers.SpacecraftHandlers, RateLimite interfaces.RateLimiter, logger interfaces.Logger)  {
 	
	gr := middleware.NewGobalRateMiddleWare(RateLimite)
	rtr:= mux.NewRouter()
	rtr.Use(gr.RateLimitMiddleware)
    rtr.Use(middleware.CORSMiddleware)

	craftRoutes := rtr.Methods(http.MethodGet, http.MethodDelete).Subrouter() 
	ga := middleware.NewAutenticateMiddleware(configApp)
    // craftRoutes.Use(middleware.AuthMiddleware )
	craftRoutes.Use(ga.AuthMiddleware)


	craftRoutes.HandleFunc("/v1/spacecrafts", spacecraftHandler.SpacecraftHandleGet).Methods(http.MethodGet)
	craftRoutes.HandleFunc("/v1/spacecrafts/{id}",spacecraftHandler.SpacecraftHandleGetByID).Methods(http.MethodGet)
	// rtr.HandleFunc("/v1/spacecrafts", spacecraftHandler.c).Methods(http.MethodPost)
	// rtr.HandleFunc("/v1/spacecrafts/", spacecraftHandler(repo, logger))
	// rtr.HandleFunc("/v1/spacecrafts/", spacecraftHandler(repo, logger))

	http.Handle("/", rtr) 
	logger.Info("Router initialized successfully on :8080")
    // log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(configApp.Server.Addr, nil))
}

 