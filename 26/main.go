package main

import (
	"fmt"
	"strings"
)

func checkUniqChar(s string) bool {
	set := map[rune]struct{}{}
	for _, val := range strings.ToLower(s) {
		if _, ok := set[val]; ok {
			return false
		}
		set[val] = struct{}{}
	}
	return true
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(checkUniqChar(s))
}
