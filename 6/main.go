package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	quit := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("Завершение через канал")
				return
			default:
				time.Sleep(2 * time.Second)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Завершение через контекст")
				return
			default:
				time.Sleep(2 * time.Second)
			}
		}
	}()

	go func() {
		fmt.Println("Завершение через Goexit()")
		runtime.Goexit()
		fmt.Println("Эта строка не выведится")
	}()

	go func() {
		for i := 0; i < 10; i++ {
			if i == 5 {
				fmt.Println("Завершение через условие")
			}
		}
	}()

	close(quit)
	cancel()
	time.Sleep(3 * time.Second)
}
