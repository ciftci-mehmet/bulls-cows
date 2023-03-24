package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	randomOptions = "0123456789abcdefghijklmnopqrstuvwxyz"
)

func main() {
	// get length input and validate
	length, err := getLength()
	if err != nil {
		fmt.Println(err)
		return
	}

	// get number of symbols input and validate
	numberOfSymbols, err := getNumOfSym(length)
	if err != nil {
		fmt.Println(err)
		return
	}

	// generate secret
	secret := GenerateSecret(length, numberOfSymbols, randomOptions)
	fmt.Printf("The secret code is prepared: %s %s.\n", GenerateStars(length), usedCharsInSecret(numberOfSymbols))

	// fmt.Println("DEBUG: secret =", secret)

	// start game loop
	fmt.Println("Okey, let's start a game!")

	turn := 1
	for {
		fmt.Printf("Turn %d. Answer:\n", turn)

		// get guess input and validate
		guess, err := getGuess(length, numberOfSymbols, randomOptions)
		if err != nil {
			fmt.Println(err)
			return
		}

		// print bulls and cows
		bulls, cows := BullsCows(secret, guess)
		PrintBullsCows(bulls, cows)

		// win condition
		if length == bulls {
			fmt.Println("Congratulations! You guessed the secret code.")
			return
		}

		turn++
	}
}

func GenerateSecret(length, numberOfSymbols int, options string) string {
	var secret string
	secretOptions := options[:numberOfSymbols]

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		randomKey := rand.Intn(numberOfSymbols)
		secret += string(secretOptions[randomKey])
	}

	return secret
}

func PrintBullsCows(bulls, cows int) {
	str := "Grade: "

	if bulls == 0 && cows == 0 {
		str += "None"
	}

	if bulls > 0 {
		str += strconv.Itoa(bulls) + " bull"
	}
	if bulls > 1 {
		str += "s"
	}

	if bulls > 0 && cows > 0 {
		str += " "
	}

	if cows > 0 {
		str += strconv.Itoa(cows) + " cow"
	}
	if cows > 1 {
		str += "s"
	}

	str += "."

	fmt.Printf("%s\n", str)
}

func BullsCows(secret, guess string) (int, int) {
	var bulls, cows int

	bsSecret := []byte(secret)
	bsGuess := []byte(guess)

	length := len(secret)

	//check bulls
	for i := 0; i < length; i++ {
		if bsGuess[i] == bsSecret[i] {
			bulls++
			bsSecret[i] = '*'
			bsGuess[i] = '-'
		}
	}

	//check cows
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if bsGuess[i] == bsSecret[j] {
				cows++
				bsSecret[j] = '*'
				bsGuess[i] = '-'
			}
		}
	}

	return bulls, cows
}

func usedCharsInSecret(numberOfSymbols int) string {
	str := "("

	//numbers
	if numberOfSymbols > 0 {
		str += string(randomOptions[0])
		if numberOfSymbols > 1 {
			str += "-"
			if numberOfSymbols > 9 {
				str += string(randomOptions[9])
			} else {
				str += string(randomOptions[numberOfSymbols-1])
			}
		}
	}

	//letters
	if numberOfSymbols > 10 {
		str += ", " + string(randomOptions[10])
		if numberOfSymbols > 11 {
			str += "-" + string(randomOptions[numberOfSymbols-1])
		}
	}

	str += ")"
	return str
}
