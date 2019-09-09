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
func (f *FileSystemStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

// GetPlayerScore func
func (f *FileSystemStore) GetPlayerScore(name string) int {
	var wins int
	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}

	return wins
}

// RecordWin func
func (f *FileSystemStore) RecordWin(name string) {
	league := f.GetLeague()

	for i, player := range league {
		if player.Name == name {
			league[i].Wins++
		}
	}

	f.database.Seek(0, 0)

	json.NewEncoder(f.database).Encode(league)
}
