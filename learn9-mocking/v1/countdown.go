package main

import (
	"bytes"
	"fmt"
)

func main() {
	Countdown(&bytes.Buffer{})
}

// Countdown func
func Countdown(out *bytes.Buffer) {
	fmt.Fprint(out, "3")
}
