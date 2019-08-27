package main

import (
	"fmt"
	"net/http"
)

// PlayerServer func
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	fmt.Fprint(w, GetPlayerScore(player))
}

// GetPlayerScore func
func GetPlayerScore(name string) string {
	var score string

	switch {
	case name == "Pepper":
		score = "20"
	case name == "Floyd":
		score = "10"
	}

	return score
}
