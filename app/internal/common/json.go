package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var (
	ErrInvalidContentType = errors.New("unexpected content type")
	ErrInvalidValue       = errors.New("invalid value")
)

// Validator allows validation of request data during the unmarshalling of the request body.
type Validator interface {
	Valid() bool
}

// ReadJSON is a generic helper function which unmarshalls a request and validates it according to the rules
// set by T's Valid() method.
func ReadJSON[T Validator](r *http.Request) (T, error) {
	var v T

	contentType := r.Header.Get("Content-Type")
	if strings.ToLower(contentType) != "application/json" {
		return v, ErrInvalidContentType
	}

	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("ReadJSON: %w", err)
	}

	if !v.Valid() {
		return v, ErrInvalidValue
	}

	return v, nil
}

//func ReadJSONSimple(r *http.Request, v any) error {
//	if r.Body == nil {
//		return fmt.Errorf("missing request body")
//	}
//
//	return json.NewDecoder(r.Body).Decode(v)
//}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

//func WriteError(w http.ResponseWriter, status int, err error) {
//	err = WriteJSON(w, status, map[string]string{"error": err.Error()})
//	if err != nil {
//		panic(err)
//	}
//}
