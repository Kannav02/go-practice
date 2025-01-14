package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	spanish            = "Spanish"
	french             = "French"
)

func HelloWorld(name string, language string) string {
	if name == "" {
		name = "World"
	}
	
	return greetingLanguage(language) + name

}

func greetingLanguage(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix

	}

	return

}

func main() {
	fmt.Println(HelloWorld("World", ""))
}
