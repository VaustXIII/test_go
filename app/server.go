package app

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

var data *LeaderBoard

func Init() {
	data = NewLeaderBoard()
}

func Run() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, getHello())
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
