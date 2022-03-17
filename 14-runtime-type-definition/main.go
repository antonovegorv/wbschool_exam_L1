// Task 14
// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.

package main

import "fmt"

type foo struct{}

func main() {
	var x interface{} = foo{}

	switch x.(type) {
	case int:
		fmt.Println("It is an Integer!")
	case string:
		fmt.Println("It is a String!")
	case foo:
		fmt.Println("It is a foo struct!")
	default:
		fmt.Println("I don't know what it is...")
	}
}
