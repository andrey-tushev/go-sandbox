package main

import (
	"fmt"
	"time"

	"net/http"
	_ "net/http/pprof"
)

func main() {
	var i int
	go func() {
		for {
			i++
			time.Sleep(3 * time.Millisecond)
		}
	}()

	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	time.Sleep(60 * time.Second)
	fmt.Println("Result", i)
}
