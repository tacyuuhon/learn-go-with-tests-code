package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Countdown(os.Stdout)
}

// Countdown func
func Countdown(out io.Writer) {
	fmt.Fprint(out, `3
	2
	1
	Go!`)
}
