package poker

import "time"

// Game struct
type Game struct {
	alerter BlindAlerter
	store   PlayerStore
}

// Start func
func (p *Game) Start(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

// Finish func
func (p *Game) Finish(winner string) {
	p.store.RecordWin(winner)
}
