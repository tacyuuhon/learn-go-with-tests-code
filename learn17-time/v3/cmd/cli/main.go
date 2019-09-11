package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/tacyuuhon/learn-go-with-tests-code/learn17-time/v3"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	store, err := poker.FileSystemStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}

	game := poker.NewCLI(store, os.Stdin, nil)
	game.PlayPoker()
}
