// Task 17
// Реализовать бинарный поиск встроенными методами языка.

package main

import "fmt"

func main() {
	data := []int{-20, -10, 0, 12, 25, 33}
	fmt.Println(binarySearch(data, 1))
}

func binarySearch(data []int, el int) int {
	mid := len(data) / 2
	switch {
	case len(data) == 0:
		return -1
	case data[mid] > el:
		return binarySearch(data[:mid], el)
	case data[mid] < el:
		var result int
		if result = binarySearch(data[mid+1:], el); result >= 0 {
			result += mid + 1
		}
		return result
	default:
		return mid
	}
}
