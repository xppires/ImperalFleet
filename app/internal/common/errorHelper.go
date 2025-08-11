package common

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandleErrorMsg  is a utility function to handle errors in a consistent way
func HandleErrorMsg(w http.ResponseWriter, msg string, statusCode int) {

	HandleError(w, nil, statusCode, msg)

}

// HandleErrorSimple HandleError is a utility function to handle errors in a consistent way
func HandleErrorSimple(w http.ResponseWriter, err error, statusCode int) {
	if err != nil {
		HandleError(w, err, statusCode, fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	HandleError(w, nil, statusCode, "Error: no error provided")

}

// HandleError HandleSuccess is a utility function to handle successful responses
func HandleError(w http.ResponseWriter, data interface{}, statusCode int, message string) {
	w.Header().Add("Content-Type", "application/json")
	// If data is nil, just send an empty response with the status code
	if data == nil {
		w.WriteHeader(statusCode)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    statusCode,
			"message": message,
			"body":    "no content",
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("Error encoding response: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		return
	}
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    statusCode,
		"message": message,
		"body":    data,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %s", err.Error()), http.StatusInternalServerError)
		return
	}

}
