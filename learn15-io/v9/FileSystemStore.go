package main

import (
	"encoding/json"
	"os"
)

// FileSystemStore struct
type FileSystemStore struct {
	database *json.Encoder
	league   League
}

// NewFileSystemStore func
func NewFileSystemStore(file *os.File) *FileSystemStore {
	file.Seek(0, 0)
	league, _ := NewLeague(file)

	return &FileSystemStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}
}

// GetLeague func of FileSystemStore
func (f *FileSystemStore) GetLeague() League {
	return f.league
}

// GetPlayerScore func
func (f *FileSystemStore) GetPlayerScore(name string) int {
	var wins int
	player := f.league.Find(name)

	if player != nil {
		wins = player.Wins
	}

	return wins
}

// RecordWin func
func (f *FileSystemStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.database.Encode(f.league)
}
