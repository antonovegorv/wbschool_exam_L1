// Task 02
// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create array.
	array := [...]int{2, 4, 6, 8, 10}

	wg := &sync.WaitGroup{} // Create WaitGroup
	wg.Add(len(array))      // Increase WaitGroup counter by array length

	// Iterate through array.
	for _, val := range array {
		// Start goroutine.
		go square(wg, val)
	}

	// Wait all goroutines to stop execution.
	wg.Wait()
}

func square(wg *sync.WaitGroup, num int) {
	defer wg.Done() // Decrease WaitGroup counter.
	fmt.Printf("Square of %v is %v\n", num, num*num)
}
