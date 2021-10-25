package common

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ResponseError struct {
	Error string `json:"error,omitempty"`
}

var (
	BadRequest       = errors.New("Bad request")
	ServerError      = errors.New("Server error")
	UserAlreadyExist = errors.New("User with that email already exists")
	UserNotExist     = errors.New("User with that id does not exist")
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case BadRequest:
		return http.StatusBadRequest
	case ServerError:
		return http.StatusInternalServerError
	case UserAlreadyExist:
		return http.StatusForbidden
	case UserNotExist:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if payload != nil {
		response, _ := json.Marshal(payload)
		_, _ = w.Write(response)
	}
}
