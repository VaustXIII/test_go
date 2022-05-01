package handlers

import models "github.com/VaustXIII/test_go/app/models"

var leaderboard *models.Leaderboard

func Initialize(lb *models.Leaderboard) {
	leaderboard = lb
}
