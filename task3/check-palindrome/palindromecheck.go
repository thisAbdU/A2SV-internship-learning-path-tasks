package palindrome

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func IsPalindrome(word string) bool {
	if len(word) == 0{
		return true
	}
	letters := []rune{}
	for _, char := range word {
		letters = append(letters, char)
	}

	i := 0
	j := len(letters) - 1

	for i < j {
		if letters[i] != letters[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func CheckPalindrome(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a word to check if it's a palindrome: ")
	word, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	word = strings.TrimSpace(word)

	if word == "" {
		fmt.Println("No word provided.")
		os.Exit(1)
	}

	if IsPalindrome(word) {
		fmt.Printf("The word '%s' is a palindrome.\n", word)
	} else {
		fmt.Printf("The word '%s' is not a palindrome.\n", word)
	}
}

