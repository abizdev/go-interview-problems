package main

import (
	cuncurrentQueue "go-problems/cuncurrent-queue"
	"sync"
)

func main() {
	// merge.MergeChannels()

	wg := &sync.WaitGroup{}
	queue := cuncurrentQueue.NewConcurrentQueue()

	wg.Add(1)
	go func() {
		defer wg.Done()
		queue.Enqueue(1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		queue.Dequeue()
	}()

	go func() {
		wg.Wait()
		queue.Len()
	}()

}
