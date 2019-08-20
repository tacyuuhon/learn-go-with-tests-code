package main

import "fmt"

//Hello func
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("World!"))
}
