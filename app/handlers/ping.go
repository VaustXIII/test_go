package handlers

import (
	"log"
	"net/http"
)

var Ping = basicHandlerWrapper(handlePing)

func handlePing(responseWriter http.ResponseWriter, request *http.Request) {
	log.Printf("Handling %s request\n%+v", request.URL, *request)
	WriteResponse(responseWriter, http.StatusOK, nil)
}
