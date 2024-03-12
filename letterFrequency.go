package main

func letterFrequency(word string) map[rune]int {
	wordFreq := make(map[rune]int)
	for _, char := range word {
		wordFreq[char]++
	}
	return wordFreq
}
