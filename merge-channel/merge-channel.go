package merge

import (
	"errors"
	"fmt"
	"sync"
)

func merge(channels ...<-chan int) (<-chan int, error) {
	if len(channels) == 0 {
		return nil, errors.New("channels are empty")
	}

	wg := sync.WaitGroup{}
	wg.Add(len(channels))

	outChannel := make(chan int)

	for _, channel := range channels {
		go func() {
			defer wg.Done()

			for value := range channel {
				outChannel <- value
			}
		}()
	}

	go func() {
		wg.Wait()
		close(outChannel)
	}()

	return outChannel, nil
}

func MergeChannels() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	chan3 := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {

			chan1 <- i
		}

		close(chan1)
	}()

	go func() {
		for i := 100; i < 200; i++ {
			chan2 <- i
		}

		close(chan2)
	}()

	go func() {
		for i := 200; i < 300; i++ {
			chan3 <- i
		}

		close(chan3)
	}()

	channels, mergeError := merge(chan1, chan2, chan3)

	if mergeError != nil {
		fmt.Println(mergeError)
	}

	go func() {
		for value := range channels {
			fmt.Println(value)
		}
	}()
}
