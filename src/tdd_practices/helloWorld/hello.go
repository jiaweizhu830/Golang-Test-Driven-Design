// Need to have main package defined with a main func inside it
package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World";
	}

	return greetingPrefix(language) + name
}

// named return value: prefix => create a variable named prefix and assigned with ""
func greetingPrefix(language string) (prefix string) {
	switch language {
		case french: 
			prefix = frenchHelloPrefix
		case spanish:
			prefix = spanishHelloPrefix
		default:
			prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}

/**
 private function name: starts with lower case
 public function name: starts with CAPITAL case
*/