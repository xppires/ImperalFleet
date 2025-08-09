package handlers

import (
	"net/http"
	"strconv"
	"app/internal/services"
	"app/internal/common"
	"app/internal/models"  
	"errors" 
	"github.com/gorilla/mux"
	"log"
)
type SpacecraftHandlers struct {
	spacecraftService *services.SpacecraftService 	
}

func NewSpacecrafHandlers(spacecraftService *services.SpacecraftService) *SpacecraftHandlers {
	return &SpacecraftHandlers{
		spacecraftService: spacecraftService, 
	}
}

var (
	errNotFound = errors.New("entry not found")
)

func (s *SpacecraftHandlers) RegisteRoutes(router *mux.Router) {
	router.HandleFunc("/spacecrafts", s.get).Methods(http.MethodGet)
	router.HandleFunc("/spacecrafts/{id}",s.GetByID).Methods(http.MethodGet)
	router.HandleFunc("/spacecrafts", s.Create).Methods(http.MethodPost)
	router.HandleFunc("/spacecrafts/{id}",s.Update).Methods(http.MethodPut)
	// router.HandleFunc("/spacecrafts/{id}",s.Delete).Methods(http.MethodDelete)
	router.Handle("/spacecrafts/{id}",http.HandlerFunc(s.Delete)).Methods(http.MethodDelete)
}

func (s *SpacecraftHandlers) Create(w http.ResponseWriter, r *http.Request) {
	
	craft, err := common.ReadJSON[*models.SpacecraftRequest](r)
	if err != nil {
		common.HandleError( w, err, http.StatusBadRequest, "invalid or malformed json"+err.Error())
		return
	}

	ctx := r.Context()		
	_,err = s.spacecraftService.Create(ctx,craft)
	if err != nil {
		common.HandleError( w, err, http.StatusInternalServerError, "failed to create entry"+err.Error())
		return
	}

	
	if err := common.WriteJSON(w, http.StatusCreated, map[string]interface{}{
			"code":   "200",
			"message": "success",
			"body":    "created",
		}); err != nil {
		common.HandleError( w, err, http.StatusInternalServerError, "failed to provide response")
		return
	}
	
}

func (s *SpacecraftHandlers) Update(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r) 
    Id, ok := pathVars["id"]
	if !ok {
		common.HandleError( w, nil, http.StatusBadRequest, "invalid or request")
		return
	}

	craft, err := common.ReadJSON[*models.SpacecraftRequest](r)
	if err != nil {
		common.HandleError( w, err, http.StatusBadRequest, "invalid or malformed json"+err.Error())
		return
	}

	ctx := r.Context()		
	err = s.spacecraftService.Update(ctx,Id,craft)
	if err != nil {
		common.HandleError( w, err, http.StatusInternalServerError, "failed to update entry "+err.Error())
		return
	}

	
	if err := common.WriteJSON(w, http.StatusOK, map[string]interface{}{
			"code":   "200",
			"message": "success",
			"body":    "updated",
		}); err != nil {
		common.HandleError( w, err, http.StatusInternalServerError, "failed to provide response")
		return
	}
	
}


func (s *SpacecraftHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	type response struct {
		success string `json:"success"`
	}
	
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		common.HandleError( w, err, http.StatusBadRequest, "invalid spaceship id")
		return
	}

	ctx := r.Context()		
	err = s.spacecraftService.Delete(ctx,id)
	if err != nil {
		switch err {
		case errNotFound:
			common.HandleError( w, err, http.StatusNotFound, "entry not found")
		default:
			common.HandleError( w, err, http.StatusInternalServerError, "failed to delete entry"+err.Error())
		}
		return
	}

	if err := common.WriteJSON(w, http.StatusCreated, map[string]interface{}{
			"code":   "200",
			"message": "success",
			"body":    "deleted",
		}); err != nil {
		common.HandleError( w, err, http.StatusInternalServerError, "failed to provide response")
		return
	}
	
}

func (s *SpacecraftHandlers) GetByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		common.HandleError( w, err, http.StatusBadRequest, "invalid id"+err.Error())
		return
	}

	ctx := r.Context()	
	spacecraft, err := s.spacecraftService.GetByID(ctx, id, nil)
	if err != nil {
		switch err {
		case errNotFound:
			common.HandleError( w, err, http.StatusNotFound, "entry not found")
		default:
			common.HandleError( w, err, http.StatusInternalServerError, "failed to retrieve entry"+err.Error())
		}
		return
	}

	if err := common.WriteJSON(w, http.StatusOK, spacecraft); err != nil {
		common.HandleError( w, err, http.StatusInternalServerError, "failed to produce response")
		return
	}
	
}

// filter by name, class, status
func (s *SpacecraftHandlers) get(w http.ResponseWriter, r *http.Request) {
	
	log.Println("Get called")
	q := r.URL.Query()
	values := make(map[string][]string)
	for key, value := range q {
		values[key] = value
	}
	// if len(values) == 0 {
	// 	common.HandleErrorMsg(w, "No filters provided", http.StatusBadRequest)
	// 	return
	// }	
	ctx := r.Context()		

	spacecrafts, err := s.spacecraftService.Get(ctx,&values)
	if err != nil {
		common.HandleError( w, err.Error(), http.StatusInternalServerError, "failed to retrieve spaceships")
		return
	}

	if err := common.WriteJSON(w, http.StatusOK, spacecrafts); err != nil {
		common.HandleError( w, err, http.StatusInternalServerError, "failed to produce response")
		return
	}

}