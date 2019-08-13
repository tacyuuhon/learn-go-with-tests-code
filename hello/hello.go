package main

import (
	"fmt"
)

const helloPrefix = "Hello, "

func hello(name string) string {
	return helloPrefix + name
}

func main() {
	fmt.Println(hello("World!"))
}
