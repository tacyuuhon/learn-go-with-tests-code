package main

import (
	"log"
	"net/http"
	"os"

	poker "github.com/tacyuuhon/learn-go-with-tests-code/learn18-websockets/v4"
)

const dbFileName = "game.db.json"
const templatePath = "../../templates/"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	server := poker.NewPlayerServer(store, templatePath)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
