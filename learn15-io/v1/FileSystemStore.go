package main

import (
	"encoding/json"
	"io"
)

// FileSystemStore struct
type FileSystemStore struct {
	database io.Reader
}

// GetLeague func of FileSystemStore
func (f *FileSystemStore) GetLeague() []Player {
	var league []Player
	json.NewDecoder(f.database).Decode(&league)
	return league
}
