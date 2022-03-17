// Task 03
// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
// квадратов(22+32+42….) с использованием конкурентных вычислений.

package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create array.
	array := [...]int{2, 4, 6, 8, 10}

	// Create a variable to store the result.
	arraySum := 0

	// Create WaitGroup.
	wg := &sync.WaitGroup{}
	wg.Add(len(array))

	// Create mutex to lock the result variable.
	mu := &sync.Mutex{}

	// Iterate through array.
	for _, val := range array {
		// Start goroutine.
		go square(wg, mu, &arraySum, val)
	}

	// Wait for all goroutines to stop execution.
	wg.Wait()

	fmt.Printf("Squares sum is %v\n", arraySum)
}

func square(wg *sync.WaitGroup, mu *sync.Mutex, sum *int, num int) {
	defer wg.Done() // Decrease WaitGroup counter

	mu.Lock()         // Lock variable
	defer mu.Unlock() // Unlock it

	*sum += num * num // Calculate square and add it to sum
}
