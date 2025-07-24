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