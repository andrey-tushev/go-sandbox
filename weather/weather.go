package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func getWeather() int {
	time.Sleep(5 * time.Second)
	return rand.Intn(70) - 30
}

// Обновление в фоне по таймеру

func main() {
	curTemp := 0
	mu := sync.RWMutex{}

	go func() {
		ticker := time.NewTicker(3000 * time.Millisecond)
		for {
			<-ticker.C
			t := getWeather()
			fmt.Println("Updated temp", t)

			mu.Lock()
			curTemp = t
			mu.Unlock()
		}
	}()

	http.HandleFunc("/weather", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Println("Request", req.URL)

		mu.RLock()
		t := curTemp
		mu.RUnlock()

		if t == 0 {
			resp.WriteHeader(500)
			fmt.Fprintf(resp, "Come back later\n")
		} else {
			fmt.Fprintf(resp, "The weather is %d\n", t) // в сервере закрывать не надо
		}
	})

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

// Кэширование всей ручки
/*
func main() {
	curTemp := 0
	mu := sync.Mutex{}
	var last int64

	http.HandleFunc("/weather", func(resp http.ResponseWriter, req *http.Request) {
		mu.Lock()

		now := time.Now().Unix()
		if now-last > 10 {
			fmt.Println("Updating")
			curTemp = getWeather()
			fmt.Println("Updated")

			last = now
		}

		fmt.Fprintf(resp, "The weather is %d\n", curTemp)
		mu.Unlock()
	})

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
*/
