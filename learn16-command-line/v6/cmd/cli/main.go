package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/tacyuuhon/learn-go-with-tests-code/learn16-command-line/v6"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}

	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
