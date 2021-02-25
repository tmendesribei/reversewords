package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text := getText()
	words := getWordsFromText(text)
	printInverseWords(words)
}

func getText() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the sentence: ")
	text, _ := reader.ReadString('\n')
	return strings.TrimSuffix(text, "\n")
}

func getWordsFromText(text string) []string {
	return strings.Split(text, " ")
}

func printInverseWords(words []string) {
	for i := len(words)-1; i >= 0; i-- {
		if (strings.TrimSpace(words[i]) != "") {
			fmt.Print(words[i])
			if i != 0 {
				fmt.Print(" ")
			}
		}
	}
	fmt.Println()
}
