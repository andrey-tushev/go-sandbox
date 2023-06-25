package main

import (
	"fmt"
	"sync"
)

func main() {
	var max int

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for i := 0; i < 1000; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()

			if i%57 == 0 {
				fmt.Println(i, i/57)

				mu.Lock()
				if i > max {
					max = i
				}
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	fmt.Println("Max is", max)
}
