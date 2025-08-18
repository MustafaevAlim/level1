package main

import "fmt"

type Caller interface {
	Call()
}

type Service1 struct {
	phone string
}

type Service2 struct {
	phone string
}

func (s *Service2) Call() {
	fmt.Println(s.phone)
}

func (s *Service1) Service1Call() string {
	return s.phone
}

type Adapter struct {
	service *Service1
}

func (a *Adapter) Call() {
	fmt.Println(a.service.Service1Call())
}

func CallToPhone(c Caller) {
	c.Call()
}

// Адптер полезен когда нужно совместить старый и новый код,создать единообразный интерфейс доступа
// для сторонних систем, совместить стороннюю библиотеку с кодом. Но паттерн усложняет архитектуру, добавляя абстракции;
// каждый вызов происходит через дополнителную прослойку, что может снизить производительность; увеличивается число
// структур.
// Допустим в приложении используется собственный интерфейс работы с БД, при подключении нового драйвера, адаптер поможет
// преобразовать вызовы собственного интерфейса к методам драйвера

func main() {
	var c Caller = &Adapter{service: &Service1{phone: "+79999999"}}
	var c1 Caller = &Service2{phone: "+78888888"}
	CallToPhone(c)
	CallToPhone(c1)
}
