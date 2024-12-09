package main

import "fmt"

type englishBot struct {
}

type spanishBot struct {
}

type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	return "Hello World"
}

func (englishBot) getGreeting1() int {
	return 10
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

/*func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}*/

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
