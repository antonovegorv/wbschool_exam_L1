// Task 04
// Реализовать постоянную запись данных в канал (главный поток). Реализовать
// набор из N воркеров, которые читают произвольные данные из канала и выводят в
// stdout. Необходима возможность выбора количества воркеров при старте.

// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ
// завершения работы всех воркеров.

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Create flag for workers.
	nWorkers := flag.Int("workers", 4, "number of workers to start")
	flag.Parse()

	// Check whether the numbers of workers greater than one.
	if *nWorkers < 1 {
		panic("there is no reason to start working with workers less than 1")
	}

	// Create channel to publish jobs.
	jobsChan := make(chan string, *nWorkers)

	// Create channel to store os signals.
	termChan := make(chan os.Signal, 1)

	// Tell program to store signals in termChan.
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	// Create WaitGroup
	wg := &sync.WaitGroup{}
	wg.Add(*nWorkers)

	// Create nWorkers.
	for i := 0; i < *nWorkers; i++ {
		// Start goroutine.
		go worker(wg, jobsChan, i)
	}

Loop:
	for i := 0; ; i++ {
		select {
		case <-termChan: // Case we got termination signal
			close(jobsChan) // Close jobsChan
			fmt.Println("Graceful shutdown executing...")
			break Loop
		default:
			// Publish a message (job) to the jobsChan.
			jobsChan <- fmt.Sprintf("Job #%v", i)
			time.Sleep(100 * time.Millisecond)
		}
	}

	// Wait all goroutines to stop execution.
	wg.Wait()
}

func worker(wg *sync.WaitGroup, jobsChan chan string, index int) {
	defer wg.Done() // Decrease WaitGroup counter

	fmt.Printf("Worker %v is starting...\n", index)
	// Range through a channel of jobs until it is closed.
	for job := range jobsChan {
		// Simulating some work in between 1 to 3 seconds.
		fmt.Printf("Worker %v has started job: %v\n", index, job)
		time.Sleep(time.Millisecond * time.Duration(1000+rand.Intn(2000)))
		fmt.Printf("Worker %v has finished processing job: %v\n", index, job)
	}
	fmt.Printf("Worker %v has been interrupted\n", index)
}
