package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

func runErrorsDemo() {
	number, err := runOrFail()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Got %v\n", number)
}

type randomError struct {
	Description string `json:"description"`
	Number      int    `json:"number"`
}

func (e *randomError) Error() string {
	jsonAsBytes, jsonMarshalError := json.Marshal(e)
	if jsonMarshalError != nil {
		return fmt.Sprintf("description: %v, number: %v", e.Description, e.Number)
	}

	return string(jsonAsBytes)
}

/* Generates a random number and will return an error if the number is negative */
func runOrFail() (int, error) {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(10) - 5

	if number < 0 {
		return 0, &randomError{"Got a negative number!!", number}
	}

	return number, nil
}
