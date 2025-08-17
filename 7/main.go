package main

import (
	"fmt"
	"math/rand"
	"sync"
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

type SyncMap struct {
	m  map[string]string
	mu sync.Mutex
}

func NewSyncMap() *SyncMap {
	return &SyncMap{m: map[string]string{}}
}

func (m *SyncMap) Push(key string, value string) {
	m.mu.Lock()
	m.m[key] = value
	m.mu.Unlock()
}

func (m *SyncMap) Delete(key string) {
	if _, ok := m.m[key]; !ok {
		return
	}
	m.mu.Lock()
	delete(m.m, key)
	m.mu.Unlock()
}

func main() {
	m := NewSyncMap()
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 5; i++ {
				m.Push(getRandomString(), getRandomString())
			}
		}()
	}
	wg.Wait()
	fmt.Println(m)

}
