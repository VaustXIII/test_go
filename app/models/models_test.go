package models

import (
	"reflect"
	"sync"
	"testing"
)

func TestLeaderboardAddAndGet(t *testing.T) {
	var tests = []struct {
		test_id  string
		adds     []Client
		expected []Client
	}{
		{"Empty", []Client{}, []Client{}},
		{"One add", []Client{
			{Id: 1, Balance: 123},
		}, []Client{
			{Id: 1, Balance: 123},
		}},
		{"Two adds", []Client{
			{Id: 1, Balance: 123},
			{Id: 5, Balance: 987.1},
		}, []Client{
			{Id: 5, Balance: 987.1},
			{Id: 1, Balance: 123},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.test_id, func(t *testing.T) {
			var leaderBoard = NewLeaderboard()
			for _, client := range tt.adds {
				leaderBoard.AddClient(client.Id, client.Balance)
			}
			var actual, readLock = leaderBoard.GetClients()
			readLock.RUnlock()

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("actual %v != expected %v", actual, tt.expected)
			}
		})
	}
}

func TestLeaderBoardGetClientBalanceNeighbours(t *testing.T) {
	var tests = []struct {
		test_id   string
		adds      []Client
		client_id int
		expected  *ClientBalanceNeighbours
	}{
		{"Empty", []Client{}, 0, nil},
		{"Only one", []Client{
			{0, 123},
		}, 0, &ClientBalanceNeighbours{Lower_id: -1, Upper_id: -1}},
		{"No upper", []Client{
			{1, 123},
			{34, 98.2},
		}, 1, &ClientBalanceNeighbours{Lower_id: 34, Upper_id: -1}},
		{"No lower", []Client{
			{2, 123},
			{43, 981.2},
		}, 2, &ClientBalanceNeighbours{Lower_id: -1, Upper_id: 43}},
		{"Basic", []Client{
			{56, 45.34},
			{43, 981.2},
			{3, 123},
		}, 3, &ClientBalanceNeighbours{Lower_id: 56, Upper_id: 43}},
		{"Longer list", []Client{
			{4, 1000},
			{10, 999},
			{11, 998},
			{12, 997},
			{13, 996},
			{14, 1001},
			{15, 1002},
			{16, 1003},
		}, 4, &ClientBalanceNeighbours{Lower_id: 10, Upper_id: 14}},
		{"Equal balances", []Client{
			{4, 1000},
			{5, 1000},
			{6, 1000},
			{12, 99},
			{13, 9999},
		}, 4, &ClientBalanceNeighbours{Lower_id: 12, Upper_id: 13}},
	}

	for _, tt := range tests {
		t.Run(tt.test_id, func(t *testing.T) {
			var leaderBoard = NewLeaderboard()
			for _, client := range tt.adds {
				leaderBoard.AddClient(client.Id, client.Balance)
			}
			var actual = leaderBoard.GetClientBalanceNeighbours(tt.client_id)

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("actual %v != expected %v", actual, tt.expected)
			}
		})
	}
}

func TestLeaderboardMultithreadSafenessAddClient(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		var leaderBoard = NewLeaderboard()

		operationsCount := 1_000

		var expected = make([]Client, operationsCount)
		for i := 0; i < operationsCount; i++ {
			var reverseId = operationsCount - 1 - i
			var reverseBalance = float32(reverseId)
			expected[i] = Client{Id: reverseId, Balance: reverseBalance}
		}

		var waitGroup sync.WaitGroup
		waitGroup.Add(3)

		go func() {
			defer waitGroup.Done()

			var innerWaitGroup sync.WaitGroup
			innerWaitGroup.Add(operationsCount)

			for i := 0; i < operationsCount; i++ {
				go func(iter int) {
					defer innerWaitGroup.Done()
					leaderBoard.AddClient(iter, float32(iter))
				}(i)
			}

			innerWaitGroup.Wait()
		}()
		go func() {
			defer waitGroup.Done()

			var innerWaitGroup sync.WaitGroup
			innerWaitGroup.Add(operationsCount)

			for i := 0; i < operationsCount; i++ {
				go func(iter int) {
					defer innerWaitGroup.Done()
					var _, readLock = leaderBoard.GetClients()
					readLock.RUnlock()
				}(i)
			}
			innerWaitGroup.Wait()
		}()

		go func() {
			defer waitGroup.Done()

			var innerWaitGroup sync.WaitGroup
			innerWaitGroup.Add(operationsCount)

			for i := 0; i < operationsCount; i++ {
				go func(iter int) {
					defer innerWaitGroup.Done()
					var _ = leaderBoard.GetClientBalanceNeighbours(iter)
				}(i)
			}
			innerWaitGroup.Wait()
		}()

		waitGroup.Wait()

		var actual, readLock = leaderBoard.GetClients()
		defer readLock.RUnlock()

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual %v != expected %v", actual, expected)
		}
	})
}
