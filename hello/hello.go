package main

import "fmt"

const englishHello = "Hello"

func Hello(name string) string {
	if len(name) == 0 {
		name = "World"
	}
	return fmt.Sprintf("%s, %s", englishHello, name)
}
func main() {
	fmt.Println(Hello("world"))
}
