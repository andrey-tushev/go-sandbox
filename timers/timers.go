package main

import (
	"fmt"
	"time"
)

func main() {
	// ticker := time.NewTicker(1 * time.Second)
	// for range ticker.C {
	// 	fmt.Println("Tick", time.Now().GoString())
	// 	time.Sleep(3 * time.Second)
	// }

	ticker1 := time.NewTicker(1 * time.Second)
	ticker2 := time.NewTicker(2 * time.Second)

	for {
		select {
		case <-ticker1.C:
			fmt.Println("ticker1")
		case <-ticker2.C:
			fmt.Println("ticker2")
		}
	}
}
