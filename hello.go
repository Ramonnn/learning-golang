package main

import "fmt"

const (
	frenchHelloPrefix  = "Bonjour, "
	spanishHelloPrefix = "Hola, "
	englishHelloPrefix = "Hello, "
	spanish            = "Spanish"
	french             = "French"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}
func greetingPrefix(language string) (prefix string) {

	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", "Spanish"))
}
