package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	messages := fmt.Sprintf(randomMessage(), name)
	return messages, nil
}

func Hellos(name []string) (map[string]string, error) {
	size := len(name)

	i := 0
	for i < size {
		if name[i] == "" {
			return nil, errors.New(string(i) + " is empty")
		}
		i++
	}

	results := make(map[string]string)
	j := 0
	for j < size {
		results[name[j]] = fmt.Sprintf(randomMessage(), name[j])
		j++
	}

	// messages := fmt.Sprintf(randomMessage(), name)
	return results, nil
}

func randomMessage() string {
	messages := []string{
		"Hi, %v. Welcome!",
		"Hi, %v. Welcome to HOme!",
		"%v. You are welcome!",
	}

	return messages[rand.Intn(len(messages))]
}
