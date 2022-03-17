// Task 07
// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	sync.RWMutex
	numbers map[int]int
}

func (sm *SafeMap) Add(num int) {
	sm.Lock()
	defer sm.Unlock()
	sm.numbers[num] = num
}

func (sm *SafeMap) Get(num int) (int, bool) {
	sm.RLock()
	defer sm.RUnlock()
	number, ok := sm.numbers[num]
	return number, ok
}

func main() {
	wg := sync.WaitGroup{}

	safeMap := &SafeMap{
		numbers: map[int]int{},
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeMap.Add(i)
		}(i)
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			number, ok := safeMap.Get(i)
			if ok {
				fmt.Println(number)
			}
		}(i)
	}

	wg.Wait()
}
