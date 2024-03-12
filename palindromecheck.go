package main

func isPalindrome(word string) bool {
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
