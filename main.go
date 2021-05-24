package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	//printGreeting(sb)
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hello!"
}

func printGreeting(sp spanishBot) {
	fmt.Println(sp.getGreeting())
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
