package greetings

/*
Export functions by beginning the function name with a capital letter.

Go executes init functions automatically at program startup, after global variables have been initialized.

Left off here: https://golang.org/doc/tutorial/greetings-multiple-people
*/

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("No name provided")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

// Seed the rand package with the current time.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns a randomly selected string from a slice.
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
