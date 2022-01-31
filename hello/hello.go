package main

import (
	"fmt"

	"github.com/baluditor/golang/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Adri")
	fmt.Println(message)
}
