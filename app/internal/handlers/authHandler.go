package handlers

import (
	"app/config"
	"app/internal/common"
	"app/internal/models"
	"app/internal/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type AuthHandler struct {
	authService *services.AuthService
	appConfig   *config.AppConfig
}

func NewAuthHandler(authService *services.AuthService, appConfig *config.AppConfig) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		appConfig:   appConfig,
	}
}

func (a *AuthHandler) RegisteRoutes(router *mux.Router) {

	router.HandleFunc("/authenticate", a.Authenticate).Methods(http.MethodPost)

}

// Authenticate handles the user authentication.
// It expects a JSON body with username and password.
// If the credentials are valid, it generates a JWT token and returns it in the response.
// If the credentials are invalid, it returns an unauthorized status.
func (a *AuthHandler) Authenticate(w http.ResponseWriter, r *http.Request) {

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid user input", http.StatusBadRequest)
		return
	}

	log.Println("authenticate called with method:", r.Method, "and :", user)

	if len(user.Username) == 0 || len(user.Password) == 0 {
		common.HandleErrorMsg(w, "Please provide username and password to obtain the token", http.StatusBadRequest)
		return
	}

	// if ok, userId, role, _:= user.Authenticate(); ok {
	if ok, userId, role, err := a.authService.Authenticate(user.Username, user.Password); ok {
		// if (user.Username == "neo" && user.Password == "keanu") || (user.Username == "morpheus" && user.Password == "lawrence") {
		if err != nil {
			common.HandleErrorMsg(w, err.Error(), http.StatusBadRequest)
			return
		}
		token, err := common.GetToken(a.appConfig.JWT.Key, userId, role)
		if err != nil {
			common.HandleErrorSimple(w, err, http.StatusInternalServerError)
		} else {
			w.Header().Add("Authorization", "Bearer "+token+userId+role)
			err := common.WriteJSON(w, http.StatusOK, map[string]interface{}{
				"code":   200,
				"userId": userId,
				"token":  token,
			})
			if err != nil {
				common.HandleErrorSimple(w, err, http.StatusInternalServerError)
			}
		}
		return
	} else {
		common.HandleErrorMsg(w, "Name and password do not match", http.StatusUnauthorized)
		return
	}
}
