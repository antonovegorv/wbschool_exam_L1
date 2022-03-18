// Task 11
// Реализовать пересечение двух неупорядоченных множеств.

package main

import (
	"fmt"
	"reflect"
)

type void struct{}

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	b := []int{2, 4, 5, 3, 8}

	fmt.Println(simpleIntersection(a, b))
}

func contains(a, e interface{}) bool {
	v := reflect.ValueOf(a)

	for i := 0; i < v.Len(); i++ {
		if v.Index(i).Interface() == e {
			return true
		}
	}
	return false
}

func simpleIntersection(a, b interface{}) []interface{} {
	set := make([]interface{}, 0)
	av := reflect.ValueOf(a)

	for i := 0; i < av.Len(); i++ {
		el := av.Index(i).Interface()
		if contains(b, el) {
			set = append(set, el)
		}
	}

	return removeDups(set)
}

func hashIntersection(a, b interface{}) []interface{} {
	set := make([]interface{}, 0)
	hash := make(map[interface{}]void)
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	for i := 0; i < av.Len(); i++ {
		el := av.Index(i).Interface()
		hash[el] = void{}
	}

	for i := 0; i < bv.Len(); i++ {
		el := bv.Index(i).Interface()
		if _, ok := hash[el]; ok {
			set = append(set, el)
		}
	}

	return removeDups(set)
}

func removeDups(a []interface{}) []interface{} {
	nodups := make([]interface{}, 0)
	encountered := make(map[interface{}]void)
	av := reflect.ValueOf(a)

	for i := 0; i < av.Len(); i++ {
		el := av.Index(i).Interface()
		if _, ok := encountered[el]; !ok {
			nodups = append(nodups, el)
			encountered[el] = void{}
		}
	}

	return nodups
}
