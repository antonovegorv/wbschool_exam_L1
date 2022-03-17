// Task 19
// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := "The quick bròwn 狐 jumped over the lazy 犬 你好好好"
	fmt.Println(reverse(s))
	fmt.Println(reverseIgnoringNonValidUTF8(s))
	fmt.Println(reversePreservingCombiningCharacters(s))
}

func reverse(s string) string {
	r := []rune(s)
	l := len(r)
	for i, j := 0, l-1; i < l/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func reverseIgnoringNonValidUTF8(s string) string {
	r := make([]rune, len(s))
	start := len(s)
	for _, c := range s {
		if c != utf8.RuneError {
			start--
			r[start] = c
		}
	}
	return string(r[start:])
}

func reversePreservingCombiningCharacters(s string) string {
	if s == "" {
		return ""
	}
	p := []rune(s)
	r := make([]rune, len(p))
	start := len(r)
	for i := 0; i < len(p); {
		if p[i] == utf8.RuneError {
			i++
			continue
		}
		j := i + 1
		for j < len(p) && (unicode.Is(unicode.Mn, p[j]) ||
			unicode.Is(unicode.Me, p[j]) || unicode.Is(unicode.Mc, p[j])) {
			j++
		}
		for k := j - 1; k >= i; k-- {
			start--
			r[start] = p[k]
		}
		i = j
	}
	return (string(r[start:]))
}
