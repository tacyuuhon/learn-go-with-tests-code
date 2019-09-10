package poker

import "io"

// CLI struct
type CLI struct {
	playerStore PlayerStore
	in          io.Reader
}

// PlayPoker func
func (cli *CLI) PlayPoker() {
	cli.playerStore.RecordWin("Panda")
}
