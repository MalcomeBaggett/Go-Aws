package main

import (
	"fmt"

	"github.com/malcomebaggett/Go-Aws/imports"
)

func main() {
	fmt.Println("Hello World")
	newTicket := imports.Ticket{
		ID:    123,
		Event: "Yoooo",
	}
	fmt.Println(newTicket)
	newTicket.PrintEvent()
}
