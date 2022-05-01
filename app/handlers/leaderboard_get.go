package handlers

import (
	"net/http"
)

var LeaderboardGet = basicHandlerWrapper(handleLeaderboardGet)

func handleLeaderboardGet(responseWriter http.ResponseWriter, request *http.Request) {
	var data = leaderboard.GetClients()

	WriteResponse(responseWriter, http.StatusOK, data)
}
