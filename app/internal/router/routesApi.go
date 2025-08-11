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

func InitRouter(configApp *config.AppConfig, spacecraftHandler *handlers.SpacecraftHandlers, RateLimite interfaces.RateLimiter, logger interfaces.Logger) {

	gr := middleware.NewGobalRateMiddleWare(RateLimite)
	rtr := mux.NewRouter()
	rtr.Use(gr.RateLimitMiddleware)
	rtr.Use(middleware.CORSMiddleware)

	craftRoutes := rtr.PathPrefix("/v1").Methods(http.MethodGet, http.MethodPut, http.MethodPost, http.MethodPut, http.MethodDelete).Subrouter()
	ga := middleware.NewAutenticateMiddleware(configApp)
	craftRoutes.Use(ga.AuthMiddleware)
	spacecraftHandler.RegisteRoutes(craftRoutes)

	http.Handle("/", craftRoutes)
	logger.Info("Router initialized successfully on :8080")
	log.Fatal(http.ListenAndServe(configApp.Server.Addr, nil))
}
