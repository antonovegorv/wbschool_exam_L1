// 23
// Удалить i-ый элемент из слайса.

package main

import "fmt"

func main() {
	// Call two methods to see the difference.
	fmt.Println(simpleRemove([]int{1, 2, 3}, 0))          // Outputs: [3 2]
	fmt.Println(orderPreservingRemove([]int{1, 2, 3}, 0)) // Outputs: [2 3]
}

// simpleRemove — removes i'th element without order preserving. Returns slice or nil if 'i'
// parameter is invalid.
func simpleRemove(s []int, i int) []int {
	// Check whether index is valid.
	if i < 0 || i > len(s) {
		return nil
	}

	// Swap last element and i'th.
	s[i] = s[len(s)-1]

	// Return slice till the last element.
	return s[:len(s)-1]
}

// orderPreservingRemove — removes i'th element with order preserving. Returns slice or nil if 'i'
// parameter is invalid.
func orderPreservingRemove(s []int, i int) []int {
	// Check whether index is valid.
	if i < 0 || i > len(s) {
		return nil
	}

	// Return slice without i'th element.
	return append(s[:i], s[i+1:]...)
}
