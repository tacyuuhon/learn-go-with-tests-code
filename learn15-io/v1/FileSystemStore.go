package main

import "io"

// FileSystemStore struct
type FileSystemStore struct {
	database io.Reader
}

// GetLeague func of FileSystemStore
func (f *FileSystemStore) GetLeague() []Player {
	league, _ := NewLeague(f.database)
	return league
}
