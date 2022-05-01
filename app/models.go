package app

type Client struct {
	Id      int
	Balance float32
}

type Leaderboard struct {
	clients map[int]Client
}

func NewLeaderboard() *Leaderboard {
	result := Leaderboard{}
	result.clients = make(map[int]Client)
	return &result
}

type ClientBalanceNeighbours struct {
	lower_id int
	upper_id int
}

func (leaderBoard *Leaderboard) AddClient(id int, balance float32) {
	leaderBoard.clients[id] = Client{Id: id, Balance: balance}
}

func (leaderBoard *Leaderboard) GetClients() *map[int]Client {
	return &leaderBoard.clients
}

func (leaderBoard *Leaderboard) GetClientBalanceNeighbours(client_id int) ClientBalanceNeighbours {
	result := ClientBalanceNeighbours{lower_id: -1, upper_id: -1}
	// TODO implement
	return result
}
