package main

import "fmt"

func BinarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	for left < right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(BinarySearch(arr, 6))
}
