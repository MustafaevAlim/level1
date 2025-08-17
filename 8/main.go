package main

import "fmt"

func setBit(num int64, i int, bit int) int64 {

	if bit == 1 {
		return num | (int64(1) << uint(i))
	} else {
		return num & ^(int64(1) << uint(i))
	}
}

func main() {
	var num int64 = 5
	var i int
	fmt.Println("Введите номер бита:")
	fmt.Scan(&i)
	i -= 1
	num = setBit(num, i, 0)
	fmt.Println(num)
}
