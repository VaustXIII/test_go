package handlers

import (
	"net/http"
)

var Ping = basicHandlerWrapper(handlePing)

func handlePing(responseWriter http.ResponseWriter, request *http.Request) {
	WriteResponse(responseWriter, http.StatusOK, nil)
}
