package main

import (
	"sync"
	"testing"
)

func testCorrectness(t *testing.T, counter Counter) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		if i%3 == 0 {
			go func(counter Counter) {
				counter.Read()
				wg.Done()
			}(counter)
		} else if i%3 == 1 {
			go func(counter Counter) {
				counter.Add(1)
				counter.Read()
				wg.Done()
			}(counter)
		} else {
			go func(counter Counter) {
				counter.Add(1)
				wg.Done()
			}(counter)
		}
	}

	wg.Wait()

	if counter.Read() != 66 {
		t.Errorf("counter should be %d and was %d", 66, counter.Read())
	}
}

func TestMutexCounter(t *testing.T) {
	testCorrectness(t, NewMutexCounter())
}

func TestChannelCounter(t *testing.T) {
	testCorrectness(t, NewChannelCounter())
}
