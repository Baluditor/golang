package greetings

import "fmt"

// Hello returns a greeting for the named person
func Hello(name string) string {
	// Returns a greeting that embed the nam in a message.
	message := fmt.Sprint("Hi, %v. Welcome", name)
	return message
}
