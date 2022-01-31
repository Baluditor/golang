package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a nem, checking
// for a valid return value
func TestHelloName(t *testing.T) {
	name := "Adri"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Adri")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Adri") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greeting.Hello with an empty string,
// check for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
