package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	arrMap := make(map[string]bool)
	for _, val := range arr {
		arrMap[val] = true
	}
	set := make([]string, 0, len(arrMap))

	for val := range arrMap {
		set = append(set, val)
	}
	fmt.Println(set)
}
