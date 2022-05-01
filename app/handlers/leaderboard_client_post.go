package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
)

var LeaderboardClientPost = basicHandlerWrapper(handleLeaderboardClientPost)

func handleLeaderboardClientPost(responseWriter http.ResponseWriter, request *http.Request) {
	type HandlerRequest struct {
		Client_id *int
		Balance   *float32
	}

	parseRequest := func(request *http.Request, parsed *HandlerRequest) error {
		var decoder = json.NewDecoder(request.Body)
		var err = decoder.Decode(&parsed)

		if parsed.Client_id == nil {
			return errors.New("client_id is a required field")
		}
		if parsed.Balance == nil {
			return errors.New("balance is a required field")
		}

		return err
	}

	var parsedRequest = HandlerRequest{}
	var parseErr = parseRequest(request, &parsedRequest)

	if parseErr != nil {
		writeErrorResponse(responseWriter, http.StatusBadRequest, parseErr.Error())
		return
	}
	leaderboard.AddClient(*parsedRequest.Client_id, *parsedRequest.Balance)

	writeResponse(responseWriter, http.StatusCreated, nil)
}
