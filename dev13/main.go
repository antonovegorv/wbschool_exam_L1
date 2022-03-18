// 13
// Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	// Create two variable that we are going to swap.
	a, b := 10, -23

	// Call for the golang-style swap.
	golangSwap(&a, &b)
	fmt.Println(a, b) // Outputs: -23 10

	// Call for the arithmetic-style swap.
	arithmeticSwap(&a, &b)
	fmt.Println(a, b) // Outputs: 10 -23

	// Call fot the bitwise-style swap.
	bitwiseSwap(&a, &b)
	fmt.Println(a, b) // Outputs -23 10
}

// golangSwap — golang-style swap.
func golangSwap(a, b *int) {
	// Check where two pointers point the same address.
	if a != b {
		*a, *b = *b, *a
	}
}

// arithmeticSwap — swaps values of two numbers via arithmetic operations (could be '/' and'*' as
// well)
func arithmeticSwap(a, b *int) {
	// Check where two pointers point the same address.
	if a != b {
		*b = *a + *b
		*a = *b - *a
		*b = *b - *a
	}
}

// bitwiseSwap — swaps values of two numbers via bitwise operations (xor).
func bitwiseSwap(a, b *int) {
	// Check where two pointers point the same address.
	if a != b {
		*a = *a ^ *b
		*b = *a ^ *b
		*a = *a ^ *b
	}
}
