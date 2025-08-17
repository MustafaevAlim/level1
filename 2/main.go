package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	wg.Add(len(nums))
	for _, num := range nums {
		go func(n int) {
			defer wg.Done()
			fmt.Fprintln(os.Stdout, n*n)
		}(num)
	}
	wg.Wait()
}
