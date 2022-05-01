package app

type Client struct {
	id      int
	balance int
}

type LeaderBoard struct {
	clients map[int]Client
}

func NewLeaderBoard() *LeaderBoard {
	result := LeaderBoard{}
	result.clients = make(map[int]Client)
	return &result
}

type ClientBalanceNeighbours struct {
	lower_id int
	upper_id int
}

func (leaderBoard *LeaderBoard) AddClient(id int, balance int) {
	leaderBoard.clients[id] = Client{id: id, balance: balance}
}

func (leaderBoard *LeaderBoard) GetClients() map[int]Client {
	return leaderBoard.clients
}

func (leaderBoard *LeaderBoard) GetClientBalanceNeighbours(client_id int) ClientBalanceNeighbours {
	result := ClientBalanceNeighbours{lower_id: -1, upper_id: -1}
	// TODO implement
	return result
}
