package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHello = "Hello"
const spanishHello = "Hola"
const frenchHello = "Bonjour"

func Hello(name string, language string) string {
	if len(name) == 0 {
		name = "World"
	}
	prefix := englishHello
	switch language {
	case spanish:
		prefix = spanishHello
	case french:
		prefix = frenchHello
	}
	return fmt.Sprintf("%s, %s", prefix, name)
}
func main() {
	fmt.Println(Hello("world", ""))
}
