package main

import "fmt"

const helloPerfix = "Hello, "

// Hello func
func Hello(name string) string {
	return helloPerfix + name
}

func main() {
	fmt.Println(Hello("World!"))
}
