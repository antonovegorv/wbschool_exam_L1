// Task 22
// Разработать программу, которая перемножает, делит, складывает, вычитает две
// числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetInt64(1<<63 - 1)
	b.SetString("9223372036854775807", 10)

	sum := new(big.Int)
	fmt.Println(sum.Add(a, b))

	sub := new(big.Int)
	fmt.Println(sub.Sub(sum, a))

	mul := new(big.Int)
	fmt.Println(mul.Mul(a, b))

	div := new(big.Int)
	fmt.Println(div.Div(mul, a))
}
