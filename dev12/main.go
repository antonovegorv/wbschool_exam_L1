// 12
// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное
// множество.

package main

import "fmt"

// Create void struct.
type void struct{}

// Create StringsSet struct.
type StringsSet struct {
	data map[string]void
}

func main() {
	// Create a slice of strings.
	s := []string{"cat", "cat", "dog", "cat", "tree"}

	// Create set from strings.
	set := stringsToSet(s)

	// Iterate through map keys and print them to the stdout.
	for k := range set.data {
		fmt.Println(k)
	}
}

// stringsToSet — converts a slice of strings to StringSet struct.
func stringsToSet(s []string) *StringsSet {
	// Initialize StringsSet struct.
	set := &StringsSet{map[string]void{}}

	// Iterate through strings.
	for _, v := range s {
		// Set value of a map with a key of string to void.
		set.data[v] = void{}
	}

	// Return the result set.
	return set
}
