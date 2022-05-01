package main

import app "github.com/VaustXIII/test_go/app"

func main() {
	var amountAddedEveryHour float32 = 10_000 // TODO: get from args/config?
	app.Run(amountAddedEveryHour)
}
