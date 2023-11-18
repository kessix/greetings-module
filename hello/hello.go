package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// set propertires of the predefined Logger
	//log.SetPrefix("greetings: ")
	//log.SetFlags(0)

	// a slice of names
	names := []string{"Kessi", "Ivone", "Diones", "Maria"}

	// request a greeting message
	//message, err := greetings.Hello("Kessix")
	messages, err := greetings.Hellos(names)

	// if an error was returned, print it to the console and exit
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returne, print the returned message
	fmt.Println(messages)
}
