package main

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("/league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)

		defer cleanDatabase()

		store := FileSystemStore{database}

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)

		defer cleanDatabase()

		store := FileSystemStore{database}

		got := store.GetPlayerScore("Chris")
		want := 33

		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)

		defer cleanDatabase()

		store := FileSystemStore{database}

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		assertScoreEquals(t, got, want)

	})
}

func assertScoreEquals(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func createTempFile(t *testing.T, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
