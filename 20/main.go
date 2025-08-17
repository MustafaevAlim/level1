package main

import (
	"fmt"
)

func reverseWordsInString(str string) string {
	s := []rune(str)
	// Сначала меняем каждые пары букв, получаем перевернутые слова
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	// Разворачиваем каждое слово обратно, чтобы получить исходное
	for i, j := 0, 0; j < len(s); {
		for j != len(s) && s[j] != ' ' {
			j++
		}

		for st := j - 1; i < st; st-- {
			s[i], s[st] = s[st], s[i]
			i++
		}
		j++
		i = j
	}

	return string(s)
}

func main() {
	fmt.Println(reverseWordsInString("snow dog sun"))
}
