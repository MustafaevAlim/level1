package main

import (
	"fmt"
	"os"
)

func reverseString(s string) string {
	str := []rune(s)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	return string(str)

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Не указана входная строка")
		return
	}

	fmt.Println(reverseString(os.Args[1]))

}
