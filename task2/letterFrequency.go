package Frequency

import (
    "strings"
)

func letterFrequency(words []string) map[string]map[rune]int {
	wordFreq := make(map[string]map[rune]int)

	for _, word := range words {
		word = normalizeWord(word)
		wordFreq[word] = make(map[rune]int)

		for _, char := range word {
			wordFreq[word][char]++
		}
	}
	return wordFreq
}

func normalizeWord(word string) string {
    // Convert the word to lowercase
    return strings.ToLower(word)
}