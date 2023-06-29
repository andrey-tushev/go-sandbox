package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Создадим несколько каналов и смержим их
	out := merge(makeChan(), makeChan(), makeChan())

	// Читаем выходной канал и печатаем значения
	fmt.Println("Значения из выходного канала")
	for v := range out {
		fmt.Println(v)
	}
}

// Функция мерджа каналов
func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}

	// Каждый входящий канал читаем из отдельных горутин и перекладываем значения в выходной канал
	for _, channel := range cs {
		wg.Add(1)
		go func(channel <-chan int) {
			for v := range channel {
				out <- v
			}
			wg.Done()
		}(channel)
	}

	// Закроем выходной канал только после того как закроются все входящие каналы
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// makeChan создает канал, кладет в него значения со случайными интервалами времени, потом закрывает канал
func makeChan() chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < 4; i++ {
			ch <- rand.Intn(1000)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(ch)
	}()

	return ch
}
