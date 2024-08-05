package Frequency

func wordFrequency(words []string) map[string]int {
	wordFreq := make(map[string]int)
	for _, word := range words {
		wordFreq[word]++
	}
	return wordFreq
}