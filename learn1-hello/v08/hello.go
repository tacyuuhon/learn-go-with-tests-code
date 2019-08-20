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

	return greetingPerfix(language) + name
}

func greetingPerfix(language string) (perfix string) {

	switch language {
	case chineseLanguage:
		perfix = chineseHelloPerfix
	case japaneseLanguage:
		perfix = japaneseHelloPerfix
	default:
		perfix = defaultHelloPerfix
	}

	return
}

func main() {
	fmt.Println(Hello("World!", ""))
}
