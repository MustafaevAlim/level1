package main

import "fmt"

func varDefinition(v interface{}) {
	fmt.Print("Тип переменной: ")
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	case chan string:
		fmt.Println("chan string")
	case chan bool:
		fmt.Println("chan bool")

	}
}

func main() {
	varDefinition("dsad")
}
