// 01
// Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в
// структуре Action от родительской структуры Human (аналог наследования).

package main

import "fmt"

// Create Human struct.
type Human struct {
	name string // Name of the Human
	age  int    // Age of the human in years
}

// sayHello — prints information about Human in stdout.
func (h *Human) sayHello() {
	fmt.Printf("Hello, my name is %v. I'm %v years old!", h.name, h.age)
}

// Create Action struct in which Human struct is embedded.
type Action struct {
	Human // All Human fields and methods become available here
}

func main() {
	// Create an instance of Action struct.
	a := Action{Human{name: "Egor Antonov", age: 21}}

	// Call for the Human method from Action struct.
	a.sayHello()
}
