package router

import (
	"log"
	"net/http" 
    "github.com/gorilla/mux"

    "app/internal/interfaces" 
	"app/internal/handlers"
	"app/internal/services"

 
)

func InitRouter(spacecraftHandler *spacecraftHandlers, logger interfaces.Logger) *http.ServeMux {
 	rtr:= mux.NewRouter()

	rtr.HandleFunc("/v1/spacecrafts", spacecraftHandler.GetAll).Methods(http.MethodGet)
	rtr.HandleFunc("/v1/spacecrafts/{id}",spacecraftHandler.GetByID).Methods(http.MethodGet)
	// rtr.HandleFunc("/v1/spacecrafts", spacecraftHandler.c).Methods(http.MethodPost)
	// rtr.HandleFunc("/v1/spacecrafts/", spacecraftHandler(repo, logger))
	// rtr.HandleFunc("/v1/spacecrafts/", spacecraftHandler(repo, logger))

	logger.Info("Router initialized successfully on :8080")
    // log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

 