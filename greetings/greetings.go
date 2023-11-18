package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// if no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}

	// OLD
	// returns a greeting that embeds the name in a message
	// message := fmt.Sprintf("Hi, %v. Welcome!", name)

	// NEW
	// create a message using random format
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil // nil (meaning no error)
}

// returns a map that associates each of the named people with a greeting message
func Hellos(names []string) (map[string]string, error) {
	// a map to associate names with message
	messages := make(map[string]string)

	//loop through the received slice of names, calling the Hello funcion to get
	// a message got each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// in the map, associate the retrieved message with the name
		messages[name] = message
	}

	return messages, nil
}

// randomFormat returns one of a set of greetings messages
// The returned message is selected at random
func randomFormat() string { // accessible only in code this package
	// a slice of message formats
	formats := []string{ // you omit its size in the brackets
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	// return a randomly selected message format by specifying
	// a random index for the slice of formats
	return formats[rand.Intn(len(formats))]
}
