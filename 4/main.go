package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func getRandomString() string {
	alphabet := []rune{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	}
	n := rand.Intn(16)
	newStr := make([]rune, n)
	for i := 0; i < n; i++ {
		newStr[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(newStr)
}

func worker(ctx context.Context, id int, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер %d завершает работу\n", id)
			return
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Канала закрыт, заверешние воркера ", id)
				return
			}
			fmt.Println(msg)

		}
	}

}

// Выбрал контексты потому что для них нужно меньше ручного контроля, чем за каналами
func main() {

	if len(os.Args) == 1 {
		fmt.Println("Не указано количество воркеров")
		return
	}

	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Неверный аргумент")
		return
	}

	var wg sync.WaitGroup
	wg.Add(N)

	ch := make(chan string, N)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer stop()

	for i := 0; i < N; i++ {
		go worker(ctx, i, ch, &wg)
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- getRandomString():

				time.Sleep(2 * time.Second)

			}

		}
	}()
	<-ctx.Done()

	fmt.Print("Завершение программы")
	wg.Wait()
	close(ch)
}
