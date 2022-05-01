package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

var LeaderboardClientNeighboursGet = basicHandlerWrapper(handleLeaderboardClientNeighboursGet)

func handleLeaderboardClientNeighboursGet(responseWriter http.ResponseWriter, request *http.Request) {
	var query = request.URL.Query()
	var client_id int
	if client_id_param, ok := query["client_id"]; ok {
		if len(client_id_param) > 1 {
			writeErrorResponse(responseWriter, http.StatusBadRequest, "Expected one client_id in query")
			return
		}

		id, err := strconv.Atoi(client_id_param[0])
		if err != nil {
			writeErrorResponse(responseWriter, http.StatusBadRequest, "client_id must be an integer")
			return
		}

		client_id = id
	} else {
		writeErrorResponse(responseWriter, http.StatusBadRequest, "Expected a client_id in query")
		return
	}

	var result = leaderboard.GetClientBalanceNeighbours(client_id)

	if result == nil {
		writeErrorResponse(responseWriter, http.StatusNotFound, fmt.Sprint("Could not find the client with id: ", client_id))
		return
	}

	writeResponse(responseWriter, http.StatusOK, result)
}
