package poker

import (
	"bufio"
	"io"
	"strings"
)

// CLI struct
type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
}

// NewCLI func
func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		in:          bufio.NewScanner(in),
	}
}

// PlayPoker func
func (cli *CLI) PlayPoker() {
	userInput := cli.readline()
	cli.playerStore.RecordWin(extractWinner(userInput))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readline() string {
	cli.in.Scan()
	return cli.in.Text()
}
