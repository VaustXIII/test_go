package app

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

const kContentType = "Content-Type"
const kApplicationJson = "application/json"

var leaderboard *Leaderboard

func Run() {
	initialize()

	http.HandleFunc("/ping", handlePing)
	http.HandleFunc("/leaderboard/client", handleLeaderboardClientPost)
	http.HandleFunc("/leaderboard", handleLeaderboardGet)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func initialize() {
	log.Println("Initializing server")
	leaderboard = NewLeaderboard()
}

func writeResponse(responseWriter http.ResponseWriter, code int, data any) {
	if data != nil {
		responseWriter.Header().Set(kContentType, kApplicationJson)
	}
	responseWriter.WriteHeader(code)
	if data != nil {
		json.NewEncoder(responseWriter).Encode(data)
	}
}

func writeErrorResponse(responseWriter http.ResponseWriter, code int, message string) {
	type ErrorResponse struct {
		Code    string
		Message string
	}

	var response = ErrorResponse{Code: http.StatusText(code), Message: message}
	writeResponse(responseWriter, code, response)
}

func handlePing(responseWriter http.ResponseWriter, request *http.Request) {
	log.Printf("Handling %s request\n%+v", request.URL, *request)
	writeResponse(responseWriter, http.StatusOK, nil)
}

func handleLeaderboardClientPost(responseWriter http.ResponseWriter, request *http.Request) {
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
		writeErrorResponse(responseWriter, http.StatusBadRequest, parseErr.Error())
		return
	}
	leaderboard.AddClient(*parsedRequest.Client_id, *parsedRequest.Balance)

	writeResponse(responseWriter, http.StatusCreated, nil)
}

func handleLeaderboardGet(responseWriter http.ResponseWriter, request *http.Request) {
	log.Printf("Handling %s request\n%+v", request.URL, *request)
	var data = leaderboard.GetClients()

	writeResponse(responseWriter, http.StatusOK, data)
}
