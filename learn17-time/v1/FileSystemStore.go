package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
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

func initialiseDBFile(file *os.File) error {
	file.Seek(0, 0)

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	return nil
}

// NewFileSystemStore func
func NewFileSystemStore(file *os.File) (*FileSystemStore, error) {
	file.Seek(0, 0)

	err := initialiseDBFile(file)

	if err != nil {
		return nil, fmt.Errorf("problem initialising db file, %v", err)
	}

	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading store from file %s, %v", file.Name(), err)
	}

	return &FileSystemStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
}

// FileSystemStoreFromFile func
func FileSystemStoreFromFile(path string) (*FileSystemStore, error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", path, err)
	}

	store, err := NewFileSystemStore(db)

	if err != nil {
		return nil, fmt.Errorf("problem creating file system player store, %v", err)
	}

	return store, nil
}

// GetLeague func of FileSystemStore
func (f *FileSystemStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
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
