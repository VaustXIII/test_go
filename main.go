package main

import app "github.com/VaustXIII/test_go/app"

func main() {
	// TODO get port from args/config?
	var amountAddedEveryHour float32 = 10_000 // TODO: get from args/config?
	app.Run(amountAddedEveryHour)
}
