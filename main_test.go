package main

import "testing"

func Test42(t *testing.T) {
	t.Run("Test42", func(t *testing.T) {
		ans := GetMeaningOfLifeUniverseAndStuff()
		if ans != 42 {
			t.Errorf("got %d, want %d", ans, 42)
		}
	})
}
