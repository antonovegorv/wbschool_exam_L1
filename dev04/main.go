// 04
// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
// которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора
// количества воркеров при старте.

// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы
// всех воркеров.

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
	// Create flag "-workers" to indicate number of workers to start.
	nWorkers := flag.Int("workers", 4, "number of workers to start")

	// Parse flags.
	flag.Parse()

	// Check whether we have at least one worker.
	if *nWorkers < 1 {
		panic("there is no reason to start workers pool with number of workers less than one")
	}

	// Create channel for OS signals.
	termChan := make(chan os.Signal, 1)

	// Direct SIGINT and SIGTERM signals to the termChan.
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	// Create channel for jobs messages buffered with the number of workers for the optimization.
	jobsChan := make(chan int, *nWorkers)

	// Create WaitGroup.
	wg := &sync.WaitGroup{}

	// Start N worker one by one.
	for i := 0; i < *nWorkers; i++ {
		// Increment WaitGroup counter by one.
		wg.Add(1)

		// Start worker as a goroutine.
		go worker(wg, jobsChan, i)
	}

JobsBatcher:
	for i := 0; ; i++ {
		select {
		case <-termChan:
			// Close a channel if we have received one of two termination signals.
			close(jobsChan)

			// Inform user about interrupting a program.
			fmt.Println("Interrupting program execution...")

			// Stop looping by breaking to a label.
			break JobsBatcher
		default:
			// Publish a job to the jobs channel.
			jobsChan <- i

			// Sleep some time to prevent "continuos" publishing messages.
			time.Sleep(100 * time.Millisecond)
		}
	}

	// Wait for all goroutines to complete.
	wg.Wait()
}

// worker — simulates jobs completing.
func worker(wg *sync.WaitGroup, jobsChan chan int, idx int) {
	// Decrement WaitGroup counter.
	defer wg.Done()

	// Inform user that worker has been started.
	fmt.Printf("Worker %v has been started\n", idx)

	// Iterate through jobs that has been published to the jobsChan until channel is closed.
	for job := range jobsChan {
		// Inform user about worker's stage simulate job from 1 to 3 seconds.
		fmt.Printf("Worker %v has received job #%v\n", idx, job)
		time.Sleep(time.Duration(1000+rand.Intn(2000)) * time.Millisecond)
		fmt.Printf("Worker %v has completed job #%v\n", idx, job)
	}

	// Inform user that worker has been interrupted.
	fmt.Printf("Worker %v has been interrupted\n", idx)
}
