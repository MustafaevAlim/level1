package main

import "fmt"

func main() {
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 99, 105}
	tempMap := make(map[int][]float64)
	for _, val := range temp {
		// В индексе выделяем диапазон
		tempMap[(int(val)/10)*10] = append(tempMap[(int(val)/10)*10], val)
	}
	fmt.Println(tempMap)
}
