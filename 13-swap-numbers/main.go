// Task 13
// Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	a, b := -10, 1<<63-1
	fmt.Println(a, b)

	a, b = b, a
	fmt.Println(a, b)

	arithmeticSwap(&a, &b)
	fmt.Println(a, b)

	bitwiseSwap(&a, &b)
	fmt.Println(a, b)
}

func arithmeticSwap(a, b *int) {
	if a != b {
		*b = *a + *b
		*a = *b - *a
		*b = *b - *a
	}
}

func bitwiseSwap(a, b *int) {
	if a != b {
		*a = *a ^ *b
		*b = *a ^ *b
		*a = *a ^ *b
	}
}
