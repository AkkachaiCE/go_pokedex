package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	// var cleanText []string
	words := strings.Fields(strings.TrimSpace(strings.ToLower((text))))
	// cleanText = append(cleanText, words...)
	// return cleanText
	return words
}
