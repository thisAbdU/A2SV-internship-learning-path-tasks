package main

import "fmt"

func letterFrequency(word string) map[rune]int {
	wordFreq := make(map[rune]int)
	for _, char := range word {
		wordFreq[char]++
	}
	fmt.Println(wordFreq)
	return wordFreq
}
