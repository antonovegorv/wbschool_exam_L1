// 25
// Реализовать собственную функцию sleep.

package main

import "time"

func main() {
	sleep(2 * time.Second)
}

// sleep — blocks main goroutine until time.After returns a channel.
func sleep(d time.Duration) {
	<-time.After(d)
}
