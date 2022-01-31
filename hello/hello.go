package main

import (
	"fmt"
	"log"

	"github.com/baluditor/golang/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.

	for i := 0; i < 20; i++ {
		printMessage()
	}

}

func printMessage() {
	message, err := greetings.Hello("Adri")
	// If an error was returned, print it to the console
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned print the returned
	// message to the console.

	fmt.Println(message)
}
