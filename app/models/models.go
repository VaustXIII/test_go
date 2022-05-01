package models

import "sort"

type Client struct {
	Id      int
	Balance float32
}

type Leaderboard struct {
	clients []Client
}

func NewLeaderboard() *Leaderboard {
	result := Leaderboard{}
	result.clients = make([]Client, 0)
	return &result
}

type ClientBalanceNeighbours struct {
	Lower_id int
	Upper_id int
}

func (leaderBoard *Leaderboard) AddClient(id int, balance float32) {
	leaderBoard.clients = append(leaderBoard.clients, Client{Id: id, Balance: balance})
	sort.Slice(leaderBoard.clients, func(i, j int) bool {
		return leaderBoard.clients[i].Balance > leaderBoard.clients[j].Balance // '>' instead of '<' for desc order
	})
}

func (leaderBoard *Leaderboard) GetClients() *[]Client {
	return &leaderBoard.clients
}

func (leaderBoard *Leaderboard) GetClientBalanceNeighbours(client_id int) *ClientBalanceNeighbours {
	result := ClientBalanceNeighbours{Lower_id: -1, Upper_id: -1}
	var clients = leaderBoard.clients
	var found = false
	for i := 0; i < len(clients); i++ {
		if clients[i].Id == client_id {
			if i > 0 {
				result.Upper_id = clients[i-1].Id
			}
			if i < len(clients)-1 {
				result.Lower_id = clients[i+1].Id
			}
			found = true
			break
		}
	}
	if !found {
		return nil
	}
	return &result
}