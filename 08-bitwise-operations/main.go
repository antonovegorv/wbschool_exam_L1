// Task 08
// Дана переменная int64. Разработать программу которая устанавливает i-й бит в
// 1 или 0.

package main

import "fmt"

func main() {
	fmt.Println(toggleBit(0, 63))
}

func toggleBit(n int64, i uint) (int64, bool) {
	if i > 63 {
		return n, false
	}
	return n ^ (1 << i), true
}
