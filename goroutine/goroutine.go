package main

import (
	"fmt"
)

// func main() {
// 	wg := sync.WaitGroup{}

// 	for i := 0; i < 10; i++ {
// 		j := i
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			fmt.Println("Goroutine", i, j)
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Println("Finish")
// }

func main() {
	ch := make(chan int, 10)

	for i := 0; i < 10; i++ {
		j := i
		go func() {
			fmt.Println("Goroutine", i, j)
			ch <- j
		}()
	}

	for i := 0; i < 10; i++ {
		fmt.Println("get from channel", <-ch)
	}
	fmt.Println("Finish")
}
