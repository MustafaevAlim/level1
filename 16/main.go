package main

import "fmt"

func quickSort(arr []int) []int {
	quick(arr, 0, len(arr)-1)
	return arr
}

func quick(arr []int, left, right int) {
	if left < right {
		p := partition(arr, left, right)
		fmt.Println(arr, " ", p)
		quick(arr, left, p-1)
		quick(arr, p, right)
	}
}

func partition(arr []int, left, right int) int {
	pivot := arr[(left+right)/2]
	i := left
	j := right
	for i <= j {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	return i

}

func main() {
	arr := []int{100, 0, -100, 4, 3, 50}
	arr = quickSort(arr)
	fmt.Println(arr)
}
