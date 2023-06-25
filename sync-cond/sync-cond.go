package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})

	// Many listeners
	for n := 0; n < 10; n++ {
		n := n
		go func() {
			cond.L.Lock()
			cond.Wait()

			fmt.Println("Listener", n)

			cond.L.Unlock()
		}()
	}

	// Broadcaster
	go func() {
		cond.L.Lock()

		fmt.Println("Do something")
		time.Sleep(3 * time.Second)
		fmt.Println("Done")

		cond.Broadcast()
		cond.L.Unlock()
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM)
	// Waiting for CTRL+C
	<-ch
}
