// Task 01
// Дана структура Human (с произвольным набором полей и методов). Реализовать
// встраивание методов в структуре Action от родительской структуры Human
// (аналог наследования).

package main

import "fmt"

// Create Human type.
type Human struct {
	name string
	age  int
}

// Create method for a Human type.
func (h *Human) sayHello() {
	fmt.Printf("Hello, my name is %v and I'm %v years old\n", h.name, h.age)
}

// Create Action type with Human type embedding in it.
type Action struct {
	Human
}

func main() {
	// Create instance of the Action type.
	a := Action{
		Human: Human{
			name: "Egor Antonov",
			age:  21,
		},
	}

	// Call for the sayHello() method which originally was described in Human
	// type
	a.sayHello()
}
