// Task 09
// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
// массива, во второй — результат операции x*2, после чего данные из второго
// канала должны выводиться в stdout.

package main

import "fmt"

func main() {
	array := [...]int{2, 4, 6, 8, 10}

	firstLine := make(chan int)
	secondLine := make(chan int)

	go fromArrToChan(array[:], firstLine)
	go fromChanToChan(firstLine, secondLine)

Loop:
	for {
		select {
		case val, ok := <-secondLine:
			if !ok {
				break Loop
			}
			fmt.Println(val)
		}
	}
}

func fromArrToChan(from []int, to chan<- int) {
	for _, val := range from {
		to <- val
	}
	close(to)
}

func fromChanToChan(from <-chan int, to chan<- int) {
	for {
		select {
		case val, ok := <-from:
			if !ok {
				close(to)
				return
			}
			to <- val * val
		}
	}
}
