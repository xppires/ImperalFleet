package router

import (
	"log"
	"net/http" 
    "github.com/gorilla/mux"

    "app/internal/interfaces"  
	"app/internal/handlers"

 
)

func InitRouter(spacecraftHandler *handlers.SpacecraftHandlers, logger interfaces.Logger)  {
 	rtr:= mux.NewRouter()

	rtr.HandleFunc("/v1/spacecrafts", spacecraftHandler.SpacecraftHandleGet).Methods(http.MethodGet)
	rtr.HandleFunc("/v1/spacecrafts/{id}",spacecraftHandler.SpacecraftHandleGetByID).Methods(http.MethodGet)
	// rtr.HandleFunc("/v1/spacecrafts", spacecraftHandler.c).Methods(http.MethodPost)
	// rtr.HandleFunc("/v1/spacecrafts/", spacecraftHandler(repo, logger))
	// rtr.HandleFunc("/v1/spacecrafts/", spacecraftHandler(repo, logger))

	http.Handle("/", rtr) 
	logger.Info("Router initialized successfully on :8080")
    // log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

 