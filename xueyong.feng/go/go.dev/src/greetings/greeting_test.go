package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *Testing.T) {
	name := "fenglehan"
	want := regexp.MustCompile(`b` + name + `\b`)
	message, err := Hello(name)
	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello("fenglehan") = %q, %v, want match for %#q, nil`, msg, err, want)
	}

}
