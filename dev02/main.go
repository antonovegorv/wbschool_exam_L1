// 02
// Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива
// (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"sync"
)

func main() {
	// Call for the function that concurrently calculates squares of the given array and prints them
	// to the stdout.
	printSquaresConcurrently([5]int{2, 4, 6, 8, 10})
}

// printSquaresConcurrently — concurrently calculates squares of the given array and prints them to
// the stdout.
func printSquaresConcurrently(arr [5]int) {
	// Create WaitGroup.
	wg := &sync.WaitGroup{}

	// Iterate through array.
	for _, v := range arr {
		// Increment WaitGroup counter by one.
		wg.Add(1)

		// Call for the anonymous goroutine that prints to the stdout square of the number.
		go func(num int) {
			// Decrement WaitGroup counter.
			defer wg.Done()

			// Print the square value.
			fmt.Printf("%v ", num*num)
		}(v)
	}

	// Wait for all goroutines to complete.
	wg.Wait()

	// Print new line in the end.
	fmt.Println()
}
