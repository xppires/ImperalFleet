package handlers

import (
	"net/http"
	"strconv"
	"app/internal/services"
	"app/internal/common"
	"app/internal/models"
	"app/internal/interfaces"
	"log/slog"
)
type SpacecraftHandlers struct {
	spacecraftService *services.SpacecraftService 	
}

func NewSpacecrafHandlers(spacecraftService *services.SpacecraftService, logger *slog.Logger) *SpacecraftHandlers {
	return &SpacecraftHandlers{
		spacecraftService: spacecraftService, 
	}
}


func SpacecraftHandleCreate(service *SpacecraftService) http.HandlerFunc {
	type response struct {
		success string `json:"success"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqData, err := common.ReadJSON[SpacecraftRequest](r)
		if err != nil {
			common.HandleError( w, err, http.StatusBadRequest, "invalid or malformed json")
		}

		err = service.repo.Create(ctx, reqData)
		if err != nil {
			common.HandleError( w, err, http.StatusInternalServerError, "failed to create entry")
		}

		if err := common.WriteJSON(w, r, http.StatusCreated, response{success: "true"}); err != nil {
			common.HandleError( w, err, http.StatusInternalServerError, "failed to provide response")
		}
	}
}

func SpacecraftHandleUpdate(service *SpacecraftService) http.HandlerFunc {
	type response struct {
		success string `json:"success"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			common.HandleError( w, err, http.StatusBadRequest, "invalid spaceship id")
		}

		reqData, err := common.ReadJSON[SpacecraftRequest](r)
		if err != nil {
			common.HandleError( w, err, http.StatusBadRequest, "invalid or malformed json")
		}

		err = service.repo.Update(ctx, id, reqData)
		if err != nil {
			switch err {
			case errNotFound:
				common.HandleError( w, err, http.StatusNotFound, "entry not found")
			default:
				common.HandleError( w, err, http.StatusInternalServerError, "failed to update entry")
			}
		}

		if err := common.WriteJSON(w, r, http.StatusCreated, response{success: "true"}); err != nil {
			common.HandleError( w, err, http.StatusInternalServerError, "failed to provide response")
		}
	}
}

func SpacecraftHandleDelete(service *SpacecraftService) http.HandlerFunc {
	type response struct {
		success string `json:"success"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			common.HandleError( w, err, http.StatusBadRequest, "invalid spaceship id")
		}

		err = service.repo.Delete(ctx, id)
		if err != nil {
			switch err {
			case errNotFound:
				common.HandleError( w, err, http.StatusNotFound, "entry not found")
			default:
				common.HandleError( w, err, http.StatusInternalServerError, "failed to delete entry")
			}
		}

		if err := common.WriteJSON(w, r, http.StatusCreated, response{success: "true"}); err != nil {
			common.HandleError( w, err, http.StatusInternalServerError, "failed to provide response")
		}
	}
}

func SpacecraftHandleGetByID(service *SpacecraftService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			common.HandleError( w, err, http.StatusBadRequest, "invalid id")
		}

		spacecraft, err := service.repo.GetByID(ctx, id)
		if err != nil {
			switch err {
			case errNotFound:
				common.HandleError( w, err, http.StatusNotFound, "entry not found")
			default:
				common.HandleError( w, err, http.StatusInternalServerError, "failed to retrieve entry")
			}
		}

		if err := common.WriteJSON(w, r, http.StatusOK, spacecraft); err != nil {
			common.HandleError( w, err, http.StatusInternalServerError, "failed to produce response")
		}
	}
}

// filter by name, class, status
func SpacecraftHandleGet(service *SpacecraftService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		filters := r.URL.Query()

		spacecrafts, err := service.repo.Get(ctx, filters)
		if err != nil {
			common.HandleError( w, err, http.StatusInternalServerError, "failed to retrieve spaceships")
		}

		if err := common.WriteJSON(w, r, http.StatusOK, spacecrafts); err != nil {
			common.HandleError( w, err, http.StatusInternalServerError, "failed to produce response")
		}
	}
}