package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func HandleLeaderboardClientPost(responseWriter http.ResponseWriter, request *http.Request) {
	log.Printf("Handling %s request\n%+v", request.URL, *request)
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
		WriteErrorResponse(responseWriter, http.StatusBadRequest, parseErr.Error())
		return
	}
	leaderboard.AddClient(*parsedRequest.Client_id, *parsedRequest.Balance)

	WriteResponse(responseWriter, http.StatusCreated, nil)
}
