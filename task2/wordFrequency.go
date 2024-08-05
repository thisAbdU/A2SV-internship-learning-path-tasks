package frequency

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wordFrequency(words []string) map[string]int {
	wordFreq := make(map[string]int)
	for _, word := range words {
		wordFreq[word]++
	}
	return wordFreq
}

func RunWordFrequency() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a phrase to determine the word frequency: ")
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

	frequency := wordFrequency(words)

	fmt.Println("Word frequency:")
	for word, count := range frequency {
		fmt.Printf("%s: %d\n", word, count)
	}

}