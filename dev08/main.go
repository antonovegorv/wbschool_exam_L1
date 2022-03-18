package main

import "fmt"

func main() {
	fmt.Println(toggleBit(0, 0)) // Outputs: 1 true
}

// toggleBit — toggles i'th bit by XORing n with 1 left shifted by i. Second return value signals
// about operation success.
func toggleBit(n int64, i uint) (int64, bool) {
	// Check whether i is greater than 63
	if i > 63 {
		return n, false
	}
	return n ^ (1 << i), true // Here we xor n with 1 left shifted by i
}
