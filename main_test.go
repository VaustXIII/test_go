package main

import "testing"

func Test42(t *testing.T) {
	t.Run("Test42", func(t *testing.T) {
		ans := GetHello()
		if ans != "Hi" {
			t.Errorf("got %s, want %s", ans, "Hi")
		}
	})
}
