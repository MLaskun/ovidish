package helpers

import (
	"fmt"
	"log/slog"
	"net/http"
)

func LogError(r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	slog.Error(err.Error(), "method", method, "uri", uri)
}

func ErrorResponse(w http.ResponseWriter, r *http.Request,
	status int, message any) {
	envelope := Envelope{"error": message}

	err := WriteJSON(w, status, envelope)
	if err != nil {
		LogError(r, err)
		w.WriteHeader(500)
	}
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	LogError(r, err)

	message := "server could not process your request"
	ErrorResponse(w, r, http.StatusInternalServerError, message)
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "resource could not be found"
	ErrorResponse(w, r, http.StatusNotFound, message)
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	ErrorResponse(w, r, http.StatusBadRequest, err.Error())
}

func FailedValidationResponse(w http.ResponseWriter, r *http.Request,
	errors map[string]string) {
	ErrorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("method %s is not allowed here", r.Method)
	ErrorResponse(w, r, http.StatusMethodNotAllowed, message)
}
