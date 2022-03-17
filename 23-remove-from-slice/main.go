// Task 23
// Удалить i-ый элемент из слайса.

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice)
	fmt.Println(removeV1(slice, 0))
	fmt.Println(removeV2(slice, 5))
	fmt.Println(removeV1(slice, -1))
}

func removeV1(slice []int, index int) []int {
	if index < 0 || index > len(slice) {
		return nil
	}

	return append(slice[:index], slice[index+1:]...)
}

func removeV2(slice []int, index int) []int {
	if index < 0 || index > len(slice) {
		return nil
	}

	length := len(slice)
	slice[index] = slice[length-1]
	return slice[:length-1]
}
