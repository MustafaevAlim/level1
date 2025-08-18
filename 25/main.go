package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	<-time.After(duration)

}

func testSleep() {
	fmt.Println("Начало testSleep")
	sleep(2 * time.Second)
	fmt.Println("Конец testSleep")
}

func main() {
	fmt.Println("Начало main")
	go testSleep()
	sleep(4 * time.Second)
	fmt.Println("Конец main")
}
