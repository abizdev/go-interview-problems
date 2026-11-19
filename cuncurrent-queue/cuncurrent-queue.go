package cuncurrentQueue

import (
	"fmt"
)

type ConcurrentQueue struct {
	queue chan int
}

func NewConcurrentQueue() *ConcurrentQueue {
	return &ConcurrentQueue{queue: make(chan int, 1)}
}

func (q *ConcurrentQueue) Enqueue(value int) {
	select {
	case q.queue <- value:
	default:
	}
}

func (q *ConcurrentQueue) Dequeue() (value int, ok bool) {
	select {
	case value, ok := <-q.queue:
		if !ok {
			fmt.Println(value, false)
			return value, false
		}

		fmt.Println(value, ok)
		return value, true

	default:
		fmt.Println(0, false)
		return 0, false
	}
}

func (q *ConcurrentQueue) Len() int {
	return len(q.queue)
}
