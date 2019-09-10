package poker

// CLI struct
type CLI struct {
	playerStore PlayerStore
}

// PlayPoker func
func (cli *CLI) PlayPoker() {
	cli.playerStore.RecordWin("Panda")
}
