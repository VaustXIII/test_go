package handlers

import (
	"net/http"
)

var LeaderboardGet = basicHandlerWrapper(handleLeaderboardGet)

func handleLeaderboardGet(responseWriter http.ResponseWriter, request *http.Request) {
	var data, readLock = leaderboard.GetClients()
	defer readLock.RUnlock()

	writeResponse(responseWriter, http.StatusOK, data)
}
