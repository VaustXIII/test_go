package app

import (
	"log"
	"net/http"
	"time"

	handlers "github.com/VaustXIII/test_go/app/handlers"
	models "github.com/VaustXIII/test_go/app/models"
)

var leaderboard *models.Leaderboard

func Run(amountAddedEveryHour float32) {
	initialize(amountAddedEveryHour)

	http.HandleFunc("/ping", handlers.Ping)
	http.HandleFunc("/leaderboard/client", handlers.LeaderboardClientPost)
	http.HandleFunc("/leaderboard", handlers.LeaderboardGet)
	http.HandleFunc("/leaderboard/client/neighbours", handlers.LeaderboardClientNeighboursGet)

	log.Println("Running")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func initialize(amountAddedEveryHour float32) {
	log.Println("Initializing server")
	leaderboard = models.NewLeaderboard()
	handlers.Initialize(leaderboard)

	go leaderboardUpdate(amountAddedEveryHour)
}

func leaderboardUpdate(amount float32) {
	getDurationTillNextHour := func() time.Duration {
		var truncateDuration = time.Hour
		// var truncateDuration = time.Second // for tests

		var now = time.Now()
		var currentHour = now.UTC().Truncate(truncateDuration)
		var nextHour = currentHour.Add(truncateDuration)
		var tillNextHour = nextHour.Sub(now)
		return tillNextHour
	}

	var timer = time.NewTimer(getDurationTillNextHour())
	for {
		<-timer.C
		timer.Reset(getDurationTillNextHour())

		log.Println("Updating all user balances by ", amount)
		leaderboard.AddAmountToAllClients(amount)
	}
}
