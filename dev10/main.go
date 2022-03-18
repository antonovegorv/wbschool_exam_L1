package main

import "fmt"

func main() {
	fmt.Println(groupTemperatures([]float64{-25.4, -27.0, 13.0, -16.7, 24.5, -21.0, 32.5}))
}

// groupTemperatures — groups temperature values in increments of 10 and returns a map with groups.
func groupTemperatures(t []float64) map[int][]float64 {
	// Initialize a map where we are going to save values.
	m := map[int][]float64{}

	// Iterate through temperatures.
	for _, v := range t {
		// Create a key of int.
		k := int(v/10) * 10

		// Append temperatrue value to it's group.
		m[k] = append(m[k], v)
	}

	// Return map.
	return m
}
