package di

import (
	"bytes"
	"fmt"
)

// Greet func
func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
