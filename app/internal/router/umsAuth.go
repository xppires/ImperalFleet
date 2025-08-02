package router

import (
	"log"
	"net/http" 
    "github.com/gorilla/mux"
	"app/internal/middleware"
    "app/internal/interfaces"  
	"app/internal/handlers"

 
)

func InitUmsRouter(handlers *handlers.UmsHandler, RateLimite interfaces.RateLimiter, logger interfaces.Logger)  {
 	
	gr := middleware.NewGobalRateMiddleWare(RateLimite)
	rtr:= mux.NewRouter()
	rtr.Use(gr.RateLimitMiddleware)
    rtr.Use(middleware.CORSMiddleware)

	umsRoutes := rtr.Methods(http.MethodPost, http.MethodDelete).Subrouter()  
	umsRoutes.HandleFunc("/v1/authenticate", handlers.Authenticate).Methods(http.MethodPost) 

	http.Handle("/", rtr) 
	logger.Info("Router initialized successfully on :8082")
    // log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8082", nil))
}

 