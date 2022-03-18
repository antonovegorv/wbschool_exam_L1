// Task 18
// Реализовать структуру-счетчик, которая будет инкрементироваться в
// конкурентной среде. По завершению программа должна выводить итоговое значение
// счетчика.

package main

import "sync"

type Counter interface {
	Add(int)
	Read() int
}

// --- Mutex Counter ---

type MutexCounter struct {
	mu     *sync.RWMutex
	number int
}

func NewMutexCounter() Counter {
	return &MutexCounter{
		&sync.RWMutex{},
		0,
	}
}

func (c *MutexCounter) Add(num int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.number += num
}

func (c *MutexCounter) Read() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.number
}

// --- Channel Counter ---

type ChannelCounter struct {
	ch     chan func()
	number int
}

func NewChannelCounter() Counter {
	counter := &ChannelCounter{make(chan func(), 100), 0}
	go func(counter *ChannelCounter) {
		for f := range counter.ch {
			f()
		}
	}(counter)
	return counter
}

func (c *ChannelCounter) Add(num int) {
	c.ch <- func() {
		c.number += num
	}
}

func (c *ChannelCounter) Read() int {
	ret := make(chan int)
	c.ch <- func() {
		ret <- c.number
		close(ret)
	}
	return <-ret
}
