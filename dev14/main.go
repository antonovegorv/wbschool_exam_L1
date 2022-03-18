// 14
// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool,
// channel из переменной типа interface{}.

package main

import "fmt"

func main() {
	// Create an empty interface of any value you want.
	var x interface{} = make(chan int)

	// A type assertion doesn’t really convert an interface to another data type, but it provides
	// access to an interface’s concrete value, which is typically what you want.
	// A type switch performs several type assertions in series and runs the first case with a
	// matching type.
	switch x.(type) {
	case int:
		fmt.Println("It is an Integer!")
	case string:
		fmt.Println("It is a String!")
	case bool:
		fmt.Println("It is a Boolean!")
	case chan int:
		fmt.Println("It is a Channel of Integers!")
	default:
		fmt.Printf("Type of x is %T. Value %v", x, x)
	}
}
