package main

import "fmt"

const helloPerfix = "Hello, "

// Hello func
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return helloPerfix + name
}

func main() {
	fmt.Println(Hello("World!"))
}
