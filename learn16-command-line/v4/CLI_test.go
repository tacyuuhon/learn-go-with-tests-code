package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {

	in := strings.NewReader("Chris wins\n")
	playerStore := &StubPlayerStore{}

	cli := &CLI{playerStore, in}
	cli.PlayPoker()

	want := "Panda"
	assertPlayerWin(t, playerStore, want)
}
