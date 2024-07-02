package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}
type bot interface {
	getGrerting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	
	prinGreeting(eb)
	prinGreeting(sb)
}

func prinGreeting(b bot) {
	fmt.Println(b.getGrerting())
}

func (englishBot) getGrerting() string {
	return "Hi There!"
}

func (spanishBot) getGrerting() string {
	return "Hola!"
}
