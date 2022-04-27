package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	var prefix = englishHelloPrefix
	if language == "Spanish" {
		prefix = spanishHelloPrefix
	}

	if name == "" {
		name = "World"
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("Dave", "English"))
}
