package main

import "fmt"

const chineseLanguage = "Chinese"
const japaneseLanguage = "Japanese"
const defaultHelloPerfix = "Hello, "
const chineseHelloPerfix = "你好，"
const japaneseHelloPerfix = "こんにちは、"

// Hello func
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == chineseLanguage {
		return chineseHelloPerfix + name
	}

	if language == japaneseLanguage {
		return japaneseHelloPerfix + name
	}
	return defaultHelloPerfix + name
}

func main() {
	fmt.Println(Hello("World!", ""))
}
