package main

import "fmt"

const spanish = "spanish"
const french = "french"

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "

func prefixer(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return prefix
}

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	prefix := prefixer(language)

	return prefix + name
}

func main() {
	fmt.Println(Hello("Brett", ""))
}
