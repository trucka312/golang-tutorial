package main

import (
	"errors"
	"fmt"
)

var ErrNotAllowed = errors.New("panda")

func entryInParty(age int) (string, error) {
	if age < 18 {
		return "", ErrNotAllowed
	}

	return "TICKET", nil
}

func main() {
	ticket, err := entryInParty(17)

	if err != nil {
		fmt.Println("log bug: ", err)
	}

	ticket, err = entryInParty(17)

	if err == ErrNotAllowed {
		fmt.Println("here")
	}

	fmt.Println("Ticket ", ticket)
}
