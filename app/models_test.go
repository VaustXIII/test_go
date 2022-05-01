package app

import (
	"reflect"
	"testing"
)

func TestLeaderBoardAddAndGet(t *testing.T) {

	t.Run("Empty", func(t *testing.T) {
		var leaderBoard = NewLeaderBoard()
		var expected = 0
		var actual = len(leaderBoard.GetClients())
		if actual != expected {
			t.Errorf("got %d, want %d", actual, expected)
		}
	})

	t.Run("Adding a client", func(t *testing.T) {
		var leaderBoard = NewLeaderBoard()
		leaderBoard.AddClient(1, 123)

		var expected = map[int]Client{
			1: {id: 1, balance: 123},
		}
		var actual = leaderBoard.GetClients()
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v, want %v", actual, expected)
		}
	})

	t.Run("Adding another client", func(t *testing.T) {
		var leaderBoard = NewLeaderBoard()
		leaderBoard.AddClient(1, 123)
		leaderBoard.AddClient(5, 987.1)

		var expected = map[int]Client{
			1: {id: 1, balance: 123},
			5: {id: 5, balance: 987.1},
		}
		var actual = leaderBoard.GetClients()
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v, want %v", actual, expected)
		}
	})
}
