package frequency

import (
	"bufio"
	"fmt"
	"os"
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
    return strings.ToLower(word)
}

func RunLetterFrequency(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a phrase to determine the letter frequency: ")
	phrase, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	phrase = strings.TrimSpace(phrase)

	if phrase == "" {
		fmt.Println("No phrase provided.")
		os.Exit(1)
	}

	words := strings.Fields(phrase)

	frequency := letterFrequency(words)

	fmt.Println("Letter frequency:")
	for word, freqMap := range frequency {
		fmt.Printf("Word: %s\n", word)
		for char, count := range freqMap {
			fmt.Printf("  %c: %d\n", char, count)
		}
	}
}