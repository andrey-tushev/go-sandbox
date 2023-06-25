package main

import (
	"fmt"
	"time"
)

func a() {
	x := []int{}     // x ->	{}
	x = append(x, 0) // x ->		{0}
	x = append(x, 1) // x ->			{0, 1}
	x = append(x, 2) // x ->				{0, 1, 2, _}

	y := append(x, 3) // y -> 				{0, 1, 2, 3}
	z := append(x, 4) // z -> 				{0, 1, 3, 4}	- сюда же указывает y, поэтому последняя ячейка перетрется

	fmt.Println(y, z) // [0, 1, 2, 4], [0, 1, 2, 4]
}

func main() {
	a()
}

////////////////////////////////////////

func main() {
	timeStart := time.Now()
	_, _ = <-worker(), <-worker()
	// _ = <-worker()
	// _ = <-worker()

	println(int(time.Since(timeStart).Seconds()))
}

func worker() chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
	}()
	return ch
}
