package poker

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

const jsonContentType = "application/json"

// PlayerStore stores score information about players
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

// Player stores a name with a number of wins
type Player struct {
	Name string
	Wins int
}

// PlayerServer is a HTTP interface for player information
type PlayerServer struct {
	store PlayerStore
	http.Handler
	templatePath string
}

// NewPlayerServer creates a PlayerServer with routing configured
func NewPlayerServer(store PlayerStore, templatePath string) *PlayerServer {
	p := new(PlayerServer)

	p.store = store
	p.templatePath = templatePath

	router := http.NewServeMux()
	router.Handle("/League", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	router.Handle("/game", http.HandlerFunc(p.game))

	p.Handler = router

	return p
}

func (p *PlayerServer) game(w http.ResponseWriter, r *http.Request) {
	// tmpl, err := template.ParseFiles("../../game.html")
	tmpl, err := template.ParseFiles(p.templatePath + "/game.html")

	if err != nil {
		http.Error(w, fmt.Sprintf("problem loading temlate %s", err.Error()), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
