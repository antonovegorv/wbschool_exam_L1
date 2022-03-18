// 20
// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("snow dog sun"))
}

// reverseWords — reverses words in string.
func reverseWords(s string) string {
	// Split the string by space characters. So, we get a slice of words.
	w := strings.Fields(s)

	// Swap words.
	for i, j := 0, len(w)-1; i < len(w)/2; i, j = i+1, j-1 {
		w[i], w[j] = w[j], w[i]
	}

	// Return new string by joining words with a space.
	return strings.Join(w, " ")
}
