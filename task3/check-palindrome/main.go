package main

import (
	"flag"
	"os"
)

func main() {
	var word string

	flag.StringVar(&word, "word", "Enter a word", "The word to check if it's a palindrome")
	flag.Parse()
	if word == ""{
		os.Exit(1)
	}

	IsPalindrome(word)
}
