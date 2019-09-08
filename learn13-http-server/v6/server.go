package main

import (
	"fmt"
	"net/http"
)

// PlayerStore interface
type PlayerStore interface {
	GetPlayerScore(name string) int
}

// PlayerServer struct
type PlayerServer struct {
	store PlayerStore
}

// StubPlayerStore struct
type StubPlayerStore struct {
	scores map[string]int
}

// GetPlayerScore func
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

// PlayerServer func
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
		return
	}

	player := r.URL.Path[len("/players/"):]
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
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
