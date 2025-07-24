package handlers

import (
	"net/http"
	"strconv"
	"app/internal/services"
	"app/internal/common"
	"app/internal/models" 
	"log/slog"
	"errors" 
)
type SpacecraftHandlers struct {
	spacecraftService *services.SpacecraftService 	
}

func NewSpacecrafHandlers(spacecraftService *services.SpacecraftService, logger *slog.Logger) *SpacecraftHandlers {
	return &SpacecraftHandlers{
		spacecraftService: spacecraftService, 
	}
}


var (
	errNotFound = errors.New("entry not found")
)

func (s *SpacecraftHandlers) SpacecraftHandleCreate(w http.ResponseWriter, r *http.Request) {
	type response struct {
		success string `json:"success"`
	}


	reqData, err := common.ReadJSON[models.SpacecraftRequest](r)
	if err != nil {
		common.HandleError( w, err, http.StatusBadRequest, "invalid or malformed json")
	}

	err = s.spacecraftService.Create(reqData)
	if err != nil {
		common.HandleError( w, err, http.StatusInternalServerError, "failed to create entry")
	}

	if err := common.WriteJSON(w, http.StatusCreated, response{success: "true"}); err != nil {
		common.HandleError( w, err, http.StatusInternalServerError, "failed to provide response")
	}
	
}

func (s *SpacecraftHandlers) SpacecraftHandleUpdate(w http.ResponseWriter, r *http.Request) {
	type response struct {
		success string `json:"success"`
	}
	


	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		common.HandleError( w, err, http.StatusBadRequest, "invalid spaceship id")
	}

	reqData, err := common.ReadJSON[models.SpacecraftRequest](r)
	if err != nil {
		common.HandleError( w, err, http.StatusBadRequest, "invalid or malformed json")
	}

	err = s.spacecraftService.Update(id, reqData)
	if err != nil {
		switch err {
		case errNotFound:
			common.HandleError( w, err, http.StatusNotFound, "entry not found")
		default:
			common.HandleError( w, err, http.StatusInternalServerError, "failed to update entry")
		}
	}

	if err := common.WriteJSON(w, http.StatusCreated, response{success: "true"}); err != nil {
		common.HandleError( w, err, http.StatusInternalServerError, "failed to provide response")
	}
	
}

func (s *SpacecraftHandlers) SpacecraftHandleDelete(w http.ResponseWriter, r *http.Request) {
	type response struct {
		success string `json:"success"`
	}
	


	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		common.HandleError( w, err, http.StatusBadRequest, "invalid spaceship id")
	}

	err = s.spacecraftService.Delete(id)
	if err != nil {
		switch err {
		case errNotFound:
			common.HandleError( w, err, http.StatusNotFound, "entry not found")
		default:
			common.HandleError( w, err, http.StatusInternalServerError, "failed to delete entry")
		}
	}

	if err := common.WriteJSON(w, http.StatusCreated, response{success: "true"}); err != nil {
		common.HandleError( w, err, http.StatusInternalServerError, "failed to provide response")
	}

}

func (s *SpacecraftHandlers) SpacecraftHandleGetByID(w http.ResponseWriter, r *http.Request) {
	


	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		common.HandleError( w, err, http.StatusBadRequest, "invalid id")
	}

	spacecraft, err := s.spacecraftService.GetByID(id)
	if err != nil {
		switch err {
		case errNotFound:
			common.HandleError( w, err, http.StatusNotFound, "entry not found")
		default:
			common.HandleError( w, err, http.StatusInternalServerError, "failed to retrieve entry")
		}
	}

	if err := common.WriteJSON(w, http.StatusOK, spacecraft); err != nil {
		common.HandleError( w, err, http.StatusInternalServerError, "failed to produce response")
	}
	
}

// filter by name, class, status
func (s *SpacecraftHandlers) SpacecraftHandleGet(w http.ResponseWriter, r *http.Request) {
	

	filters := r.URL.Query()

	spacecrafts, err := s.spacecraftService.Get(filters)
	if err != nil {
		common.HandleError( w, err, http.StatusInternalServerError, "failed to retrieve spaceships")
	}

	if err := common.WriteJSON(w, http.StatusOK, spacecrafts); err != nil {
		common.HandleError( w, err, http.StatusInternalServerError, "failed to produce response")
	}

}