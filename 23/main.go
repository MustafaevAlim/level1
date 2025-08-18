package main

import "fmt"

func deleteIndex(arr []int, index int) []int {
	copy(arr[index:], arr[index+1:])
	return arr[:len(arr)-1]

}

func main() {
	arr := []int{1, 2, 6, 3, 4, 5}
	arr = deleteIndex(arr, 2)
	fmt.Println(arr)
}
