package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// PlayerPrompt const
const PlayerPrompt = "Please enter the number of players: "

// CLI struct
type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
	out         io.Writer
	game        *Game
}

// NewCLI func
func NewCLI(in io.Reader, out io.Writer, game *Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

// NewGame func
func NewGame(alerter BlindAlerter, store PlayerStore) *Game {
	return &Game{
		alerter: alerter,
		store:   store,
	}
}

// PlayPoker func
func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)

	numberOfPlayersInput := cli.readline()
	numberOfPlayers, _ := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))

	cli.game.Start(numberOfPlayers)

	winnerInput := cli.readline()
	winner := extractWinner(winnerInput)

	cli.game.Finish(winner)
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins\n", "", 1)
}

func (cli *CLI) readline() string {
	cli.in.Scan()
	return cli.in.Text()
}
