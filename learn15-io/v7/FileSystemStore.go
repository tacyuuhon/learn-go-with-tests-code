package main

import (
	"encoding/json"
	"io"
)

// FileSystemStore struct
type FileSystemStore struct {
	database io.ReadWriteSeeker
}

// GetLeague func of FileSystemStore
func (f *FileSystemStore) GetLeague() League {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

// GetPlayerScore func
func (f *FileSystemStore) GetPlayerScore(name string) int {
	var wins int
	player := f.GetLeague().Find(name)

	if player != nil {
		wins = player.Wins
	}

	return wins
}

// RecordWin func
func (f *FileSystemStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		league = append(league, Player{name, 1})
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}
