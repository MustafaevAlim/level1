package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
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

	ch := make(chan string, N)

	for i := 0; i < N; i++ {
		go func(ch chan string) {
			for msg := range ch {
				fmt.Println(msg)

			}
		}(ch)
	}

	for {
		ch <- getRandomString()
	}

}
