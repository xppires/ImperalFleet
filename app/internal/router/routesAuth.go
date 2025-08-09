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

func InitAuthRouter(appconfig *config.ConfigApp, handlers *handlers.AuthHandler, RateLimite interfaces.RateLimiter, logger interfaces.Logger)  {
 	
	gr := middleware.NewGobalRateMiddleWare(RateLimite)
	rtr:= mux.NewRouter()
	rtr.Use(gr.RateLimitMiddleware)
    rtr.Use(middleware.CORSMiddleware)

	authRoutes := rtr.PathPrefix("/v1").Methods(http.MethodPost, http.MethodDelete).Subrouter()  
	handlers.RegisteRoutes(authRoutes)
	
	http.Handle("/", rtr) 
	logger.Info("Router initialized successfully on "+appconfig.Server.Addr)
    // log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(appconfig.Server.Addr, nil))
}

 