package middleware

import (
    "app/internal/common"
	"net/http" 
	"log"
	"encoding/json" 
	"context"
)
// authenticate handles the user authentication.
// It expects a JSON body with username and password.
// If the credentials are valid, it generates a JWT token and returns it in the response.
// If the credentials are invalid, it returns an unauthorized status.
func Authenticate(w http.ResponseWriter, r *http.Request) {
	
	var user User
	 if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid user input", http.StatusBadRequest)
        return
    }
    
	log.Println("authenticate called with method:", r.Method, "and :", user)
	 
	if len(user.Username) == 0 || len(user.Password) == 0 {
		common.HandleErrorMsg(w, "Please provide username and password to obtain the token", http.StatusBadRequest)
		return
	}
	if ok, userId,role, _ := user.Authenticate(); ok { 
	// if (user.Username == "neo" && user.Password == "keanu") || (user.Username == "morpheus" && user.Password == "lawrence") {
		token, err := common.GetToken(userId, role)
		if err != nil {
			common.HandleErrorSimple(w, err, http.StatusInternalServerError)
		} else {
			w.Header().Set("Authorization", "Bearer "+token)
			w.WriteHeader(http.StatusOK)
			
			json.NewEncoder(w).Encode(map[string]interface{}{
			"code":   200,
			"userId": userId,
			"token":   token,
			}) 
			
		}
	} else {
		common.HandleErrorMsg(w, "Name and password do not match", http.StatusUnauthorized)
		return
	}
}

// authMiddleware is a middleware that checks for the JWT token in the Authorization header
// It verifies the token and extracts userId and role from the claims.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		name,role, err := common.ExtractUserFromJWT(w,r)
		if err != nil {
			common.HandleErrorSimple(w, err, http.StatusUnauthorized)
			return
		}
		
		ctx := r.Context()
		ctx = context.WithValue(ctx,"userId", name)
		ctx = context.WithValue(ctx,"Role", role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}