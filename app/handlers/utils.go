package handlers

import (
	"encoding/json"
	"net/http"
)

const kContentType = "Content-Type"
const kApplicationJson = "application/json"

func WriteResponse(responseWriter http.ResponseWriter, code int, data any) {
	if data != nil {
		responseWriter.Header().Set(kContentType, kApplicationJson)
	}
	responseWriter.WriteHeader(code)
	if data != nil {
		json.NewEncoder(responseWriter).Encode(data)
	}
}

func WriteErrorResponse(responseWriter http.ResponseWriter, code int, message string) {
	type ErrorResponse struct {
		Code    string
		Message string
	}

	var response = ErrorResponse{Code: http.StatusText(code), Message: message}
	WriteResponse(responseWriter, code, response)
}
