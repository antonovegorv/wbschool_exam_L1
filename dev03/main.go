// 03
// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием
// конкурентных вычислений.

package main

import (
	"fmt"
	"sync"
)

func main() {
	// Call for the function that concurrently calculates sum of the squares of the given array and
	// returns this value.
	fmt.Println("Sum is", calculateSquaresSumConcurrently([5]int{2, 4, 6, 8, 10}))
}

// calculateSquaresSumConcurrently — concurrently calculates sum of the squares of the given array
// and returns this value.
func calculateSquaresSumConcurrently(arr [5]int) (sum int) {
	// Create WaitGroup.
	wg := &sync.WaitGroup{}

	// Create Mutex.
	mu := &sync.Mutex{}

	// Iterate through array.
	for _, v := range arr {
		// Increment WaitGroup counter by one.
		wg.Add(1)

		// Call for the anonymous goroutine that calculates a square of the number and adds this
		// value to the result.
		go func(num int) {
			// Decrement WaitGroup counter.
			defer wg.Done()

			// Lock mutex.
			mu.Lock()

			// Unlock mutex.
			defer mu.Unlock()

			// Calculate square of the number and add this value to the result.
			sum += num * num
		}(v)
	}

	// Wait for all goroutines to complete.
	wg.Wait()

	return
}
