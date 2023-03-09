package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "fenglehan"
	want := regexp.MustCompile(`b` + name + `\b`)
	message, err := Hello(name)
	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello("fenglehan") = %q, %v, want match for %#q, nil`, message, err, want)
	}
}
