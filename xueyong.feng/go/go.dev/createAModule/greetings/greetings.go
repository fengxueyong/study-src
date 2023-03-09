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

func Hellos(names []string) (map[string]string, error) {
	size := len(names)

	// i := 0
	// for i < size {
	// 	if name[i] == "" {
	// 		return nil, errors.New(string(i) + " is empty")
	// 	}
	// 	i++
	// }
	if size < 1 {
		return nil, errors.New("names is empty")
	}

	results := make(map[string]string)

	for _, name := range names {
		_, err := Hello(name)
		if err != nil {
			return nil, errors.New(name + "fatal error")
		}

		results[name] = fmt.Sprintf(randomMessage(), name)
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
