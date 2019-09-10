package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// ReadWriteSeekTruncate interface
type ReadWriteSeekTruncate interface {
	io.ReadWriteSeeker
	Truncate(size int64) error
}

// FileSystemStore struct
type FileSystemStore struct {
	database *json.Encoder
	league   League
}

// NewFileSystemStore func
func NewFileSystemStore(file *os.File) (*FileSystemStore, error) {
	file.Seek(0, 0)

	info, err := file.Stat()

	if err != nil {
		return nil, fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	league, err := NewLeague(file)

	return &FileSystemStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
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
