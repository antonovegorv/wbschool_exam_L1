// Task 06
// Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go stopMeWithCancelContext(ctx, wg)
	time.Sleep(2 * time.Second)
	cancel()

	wg.Wait()

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	wg.Add(1)
	go stopMeWithTimeoutContext(ctx, wg)

	wg.Wait()

	ch := make(chan struct{})
	wg.Add(1)
	go stopMeWithChannelV1(ch, wg)
	time.Sleep(2 * time.Second)
	ch <- struct{}{}

	wg.Wait()

	wg.Add(1)
	go stopMeWithChannelV2(ch, wg)
	time.Sleep(2 * time.Second)
	close(ch)

	wg.Wait()

	ch = make(chan struct{})
	wg.Add(1)
	go stopMeWithChannelV3(ch, wg)
	time.Sleep(2 * time.Second)
	close(ch)

	wg.Wait()
}

func stopMeWithCancelContext(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Goroutine has started execution!")
	<-ctx.Done()
	fmt.Printf("Stopped by cancellation!\n\n")
}

func stopMeWithTimeoutContext(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Goroutine has started execution!")
	<-ctx.Done()
	fmt.Printf("Stopped by timeout!\n\n")
}

func stopMeWithChannelV1(ch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Goroutine has started execution!")
	<-ch
	fmt.Printf("Stopped by channel!\n\n")
}

func stopMeWithChannelV2(ch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Goroutine has started execution!")
	for {
		if _, ok := <-ch; !ok {
			fmt.Printf("Stopped by closing a channel!\n\n")
			return
		}
	}
}

func stopMeWithChannelV3(ch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Goroutine has started execution!")
	for range ch {
		runtime.Gosched()
	}
	fmt.Printf("Stopped by closing a channel (range)!\n\n")
}
