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

	authRoutes := rtr.Methods(http.MethodPost, http.MethodDelete).Subrouter()  
	authRoutes.HandleFunc("/v1/authenticate", handlers.Authenticate).Methods(http.MethodPost) 

	http.Handle("/", rtr) 
	logger.Info("Router initialized successfully on :8081")
    // log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(appconfig.Server.Addr, nil))
}

 