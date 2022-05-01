package handlers

import (
	"log"
	"net/http"
)

func HandleLeaderboardGet(responseWriter http.ResponseWriter, request *http.Request) {
	log.Printf("Handling %s request\n%+v", request.URL, *request)
	var data = leaderboard.GetClients()

	WriteResponse(responseWriter, http.StatusOK, data)
}
