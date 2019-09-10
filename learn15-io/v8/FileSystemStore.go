package main

import (
	"encoding/json"
	"io"
)

// FileSystemStore struct
type FileSystemStore struct {
	database io.ReadWriteSeeker
	league   League
}

// NewFileSystemStore func
func NewFileSystemStore(database io.ReadWriteSeeker) *FileSystemStore {
	database.Seek(0, 0)
	league, _ := NewLeague(database)
	return &FileSystemStore{
		database: database,
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

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(f.league)
}
