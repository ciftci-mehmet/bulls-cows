package main

import (
	"bufio"
	"os"
)

func GenerateStars(length int) string {
	var stars string
	for i := 0; i < length; i++ {
		stars += "*"
	}
	return stars
}

func requestInput() string {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	return scanner.Text()
}
