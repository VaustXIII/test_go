package app

import (
	"log"
	"net/http"

	handlers "github.com/VaustXIII/test_go/app/handlers"
	models "github.com/VaustXIII/test_go/app/models"
)

var leaderboard *models.Leaderboard

func Run() {
	initialize()

	http.HandleFunc("/ping", handlers.Ping)
	http.HandleFunc("/leaderboard/client", handlers.LeaderboardClientPost)
	http.HandleFunc("/leaderboard", handlers.LeaderboardGet)
	http.HandleFunc("/leaderboard/client/neighbours", handlers.LeaderboardClientNeighboursGet)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func initialize() {
	log.Println("Initializing server")
	leaderboard = models.NewLeaderboard()
	handlers.Initialize(leaderboard)
}
