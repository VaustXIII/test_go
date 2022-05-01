package handlers

import (
	"log"
	"net/http"
)

func HandlePing(responseWriter http.ResponseWriter, request *http.Request) {
	log.Printf("Handling %s request\n%+v", request.URL, *request)
	WriteResponse(responseWriter, http.StatusOK, nil)
}
