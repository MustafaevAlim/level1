package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func processNumber(ctx context.Context, arr []int, numbersCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(numbersCh)
	for _, v := range arr {
		select {
		case <-ctx.Done():
			fmt.Println("Завершение обработки чисел")
			return
		case numbersCh <- v:
		}
	}
}

func doubleNumber(numbersCh, doubleNumbersCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for number := range numbersCh {

		doubleNumbersCh <- number * 2
	}
	close(doubleNumbersCh)

}

func printNumbers(doubleNumbersCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for number := range doubleNumbersCh {
		fmt.Println(number)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	numbersCh := make(chan int, 10)
	doubleNumbersCh := make(chan int, 10)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT)
	defer cancel()

	arr := []int{10, 20, 30, 45, 2100}

	go processNumber(ctx, arr, numbersCh, &wg)
	go doubleNumber(numbersCh, doubleNumbersCh, &wg)
	go printNumbers(doubleNumbersCh, &wg)
	wg.Wait()
	fmt.Println("Завершение основной программы")

}
