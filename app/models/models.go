package models

import (
	"sort"
	"sync"
)

type Client struct {
	Id      int
	Balance float32 // TODO decimal?
}

type Leaderboard struct {
	clients []Client
	lock    sync.RWMutex
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
	leaderBoard.lock.Lock()
	defer leaderBoard.lock.Unlock()

	leaderBoard.clients = append(leaderBoard.clients, Client{Id: id, Balance: balance})
	sort.Slice(leaderBoard.clients, func(i, j int) bool {
		return leaderBoard.clients[i].Balance > leaderBoard.clients[j].Balance // '>' instead of '<' for desc order
	})
}

func (leaderBoard *Leaderboard) GetClients() (clients []Client, readLock *sync.RWMutex) {
	leaderBoard.lock.RLock()
	return leaderBoard.clients, &leaderBoard.lock
}

func (leaderBoard *Leaderboard) GetClientBalanceNeighbours(client_id int) *ClientBalanceNeighbours {
	result := ClientBalanceNeighbours{Lower_id: -1, Upper_id: -1}
	var clients = leaderBoard.clients
	var found = false

	leaderBoard.lock.RLock()
	defer leaderBoard.lock.RUnlock()

	for i := 0; i < len(clients); i++ {
		if clients[i].Id == client_id {
			for j := i - 1; j >= 0; j-- {
				if clients[j].Balance != clients[i].Balance {
					result.Upper_id = clients[j].Id
					break
				}
			}
			for j := i + 1; j < len(clients); j++ {
				if clients[j].Balance != clients[i].Balance {
					result.Lower_id = clients[j].Id
					break
				}
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

func (leaderBoard *Leaderboard) AddAmountToAllClients(amount float32) {
	leaderBoard.lock.Lock()
	defer leaderBoard.lock.Unlock()

	for i, _ := range leaderBoard.clients {
		leaderBoard.clients[i].Balance += amount
	}
}
