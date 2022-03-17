// Task 15
// К каким негативным последствиям может привести данный фрагмент кода, и как
//это исправить? Приведите корректный пример реализации.

package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	// justString = v[:100]
	justString = string([]byte(v[:100]))
	// justString = strings.Clone(v[:100])

	header := (*reflect.StringHeader)(unsafe.Pointer(&justString))
	fmt.Println(header.Data)

	header = (*reflect.StringHeader)(unsafe.Pointer(&v))
	fmt.Println(header.Data)
}

func main() {
	someFunc()
}

func createHugeString(size int) string {
	sb := strings.Builder{}
	sb.Grow(size)
	for i := 0; i < size; i++ {
		sb.WriteByte(48)
	}
	return sb.String()
}
