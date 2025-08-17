package main

import (
	"fmt"
	"math/rand"
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

func worker(ch <-chan string) {

	for msg := range ch {
		fmt.Println(msg)
	}

}

func main() {

	ch := make(chan string)
	go worker(ch)

	timer := time.After(5 * time.Second)
	go func(ch chan<- string) {
		for {
			select {
			case ch <- getRandomString():
				fmt.Println("Запись в канал")
			case <-timer:
				fmt.Println("Завершение работы")
				return
			}

		}
	}(ch)
	<-timer
	close(ch)
	fmt.Println("Завершение программы")

}
