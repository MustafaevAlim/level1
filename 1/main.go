package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h *Human) SetAge(age int) {
	h.age = age
}

type Action struct {
	Human
}

func (a *Action) Run() {
	fmt.Printf("Run %s %d\n", a.name, a.age)
}

func main() {
	act := Action{Human: Human{name: "Nick", age: 20}}
	act.Run()
	act.SetAge(25)
	act.Run()
}
