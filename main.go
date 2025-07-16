package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			cleanText := cleanInput(scanner.Text())
			fmt.Print("Your command was: " + cleanText[0])
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

	}
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.TrimSpace(strings.ToLower((text))))
	return words
}
