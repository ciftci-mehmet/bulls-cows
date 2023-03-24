package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func getLength() (int, error) {
	fmt.Println("Input the length of the secret code:")
	input := requestInput()

	length, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("error: \"%s\" isn't a valid number", input)
	}

	if length < 1 {
		return 0, errors.New("error: length can not be lower than 1")
	}

	return length, nil
}

func getNumOfSym(length int) (int, error) {
	fmt.Println("Input the number of possible symbols in the code:")
	input := requestInput()

	numberOfSymbols, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("error: \"%s\" isn't a valid number", input)
	}

	if numberOfSymbols < length {
		return 0, fmt.Errorf("error: it's not possible to generate a code with a length of %d with %d unique symbols", length, numberOfSymbols)
	}

	if numberOfSymbols > 36 {
		return 0, errors.New("error: maximum number of possible symbols in the code is 36 (0-9, a-z)")
	}

	return numberOfSymbols, nil
}

func getGuess(length int, numOfSym int, options string) (string, error) {
	guess := requestInput()

	if len(guess) != length {
		return "", errors.New("error: guess must be same length as secret")
	}

	availableOptions := options[:numOfSym]
	for i := 0; i < len(guess); i++ {
		if !strings.Contains(availableOptions, string(guess[i])) {
			return "", errors.New("error: guess contains invalid symbol")
		}
	}

	return guess, nil
}
