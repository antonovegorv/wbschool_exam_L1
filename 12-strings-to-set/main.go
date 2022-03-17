// Task 12
// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
// собственное множество.

package main

import "fmt"

type void struct{}

type StringsSet struct {
	m map[string]void
}

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	set := stringsToSet(strings)

	for k := range set.m {
		fmt.Println(k)
	}
}

func stringsToSet(s []string) *StringsSet {
	set := &StringsSet{map[string]void{}}
	for _, v := range s {
		set.m[v] = void{}
	}
	return set
}
