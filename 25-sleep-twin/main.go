// Task 25
// Реализовать собственную функцию sleep.

package main

import "time"

func main() {
	sleep(2 * time.Second)
}

func sleep(d time.Duration) {
	<-time.After(d)
}
