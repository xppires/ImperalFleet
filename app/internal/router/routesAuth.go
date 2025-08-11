package router

import (
	"app/config"
	"app/internal/handlers"
	"app/internal/interfaces"
	"app/internal/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func InitAuthRouter(appconfig *config.AppConfig, handlers *handlers.AuthHandler, RateLimite interfaces.RateLimiter, logger interfaces.Logger) {

	gr := middleware.NewGobalRateMiddleWare(RateLimite)
	rtr := mux.NewRouter()
	rtr.Use(gr.RateLimitMiddleware)
	rtr.Use(middleware.CORSMiddleware)

	authRoutes := rtr.PathPrefix("/v1").Methods(http.MethodPost, http.MethodDelete).Subrouter()
	handlers.RegisteRoutes(authRoutes)

	http.Handle("/", rtr)
	logger.Info("Router initialized successfully on " + appconfig.Server.Addr)
	// log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(appconfig.Server.Addr, nil))
}
