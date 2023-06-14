package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Start")

	ctx := context.Background()
	ctx, cancelFunc := context.WithTimeout(ctx, time.Second)
	defer func() {
		cancelFunc() // чтобы завершить таймер и освободить ресурсы (не забываем про это, когда работаем с контекстом)
	}()

	res, err := wrapper(ctx)
	fmt.Printf("Result: %s; Error: %v; Error type: %T\n", res, err, err)

	if err == context.DeadlineExceeded {
		fmt.Println("DeadlineExceeded!")
	}
	// if errors.Is(err, context.DeadlineExceeded) {
	// 	fmt.Println("DeadlineExceeded!")
	// }
}

func wrapper(ctx context.Context) (string, error) {
	ch := make(chan string)

	go func() {
		res := longFunc()
		ch <- res
		close(ch)
	}()

	select {
	case res := <-ch:
		fmt.Println("* ch closed")
		return res, nil

	case <-ctx.Done(): //  struct{} - чисто сигнальный канал
		fmt.Println("* ctx done")
		return "", ctx.Err() // вернет context.DeadlineExceeded
	}
}

func longFunc() string {
	d := time.Millisecond * time.Duration(rand.Intn(2000))
	time.Sleep(d)

	return "RESULT"
}
