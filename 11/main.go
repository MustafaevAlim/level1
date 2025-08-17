package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{2, 3, 4}
	arr1Map := make(map[int]int)
	for _, val := range arr1 {
		arr1Map[val]++
	}

	for _, val := range arr2 {
		if count, ok := arr1Map[val]; ok && count > 0 {
			fmt.Print(val, " ")
			arr1Map[val]--
		}
	}
	fmt.Println()

}
