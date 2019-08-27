package main

import (
	"fmt"
	"net/http"
)

// PlayerServer func
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch {
	case player == "Pepper":
		fmt.Fprint(w, "20")
	case player == "Floyd":
		fmt.Fprint(w, "10")
	}
}
