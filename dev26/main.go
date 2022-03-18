// Task 26
// Разработать программу, которая проверяет, что все символы в строке уникальные
//  (true — если уникальные, false etc). Функция проверки должна быть
// регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	str := "abcdefghijka"
	fmt.Println(areCharactersUniqueV1(str))
	fmt.Println(areCharactersUniqueV2(str))
	fmt.Println(areCharactersUniqueV3(str))
}

func areCharactersUniqueV1(str string) bool {
	r := []rune(strings.ToLower(str))
	l := len(r)

	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if r[i] == r[j] {
				return false
			}
		}
	}

	return true
}

func areCharactersUniqueV2(str string) bool {
	r := []rune(strings.ToLower(str))
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	for i := 0; i < len(r)-1; i++ {
		if r[i] == r[i+1] {
			return false
		}
	}

	return true
}

func areCharactersUniqueV3(str string) bool {
	m := make(map[rune]struct{})

	for _, ch := range str {
		if _, ok := m[ch]; ok {
			return false
		}
		m[ch] = struct{}{}
	}

	return true
}
