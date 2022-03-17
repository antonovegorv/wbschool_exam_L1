// Task 16
// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	data := []int{2, 4, 2, 1, -3, 4, 9, 1}
	fmt.Println(quickSort(data))
}

func quickSort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	left, right := 0, len(data)-1
	pivot := rand.Int() % len(data)

	data[pivot], data[right] = data[right], data[pivot]

	for i := range data {
		if data[i] < data[right] {
			data[left], data[i] = data[i], data[left]
			left++
		}
	}

	data[left], data[right] = data[right], data[left]

	quickSort(data[:left])
	quickSort(data[left+1:])

	return data
}
