package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	m     sync.Mutex
}

func (c *Counter) Inc() {
	c.m.Lock()
	c.count++
	c.m.Unlock()
}

func (c *Counter) Get() int {
	return c.count
}

func main() {
	counter := &Counter{}
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(c *Counter) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				c.Inc()
			}
		}(counter)
	}
	wg.Wait()
	fmt.Println(counter.Get())
}
