package main

import "fmt"

import "rsc.io/quote"

import "example.com/greetings"

func main() {
	fmt.Println(quote.Go())

	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
