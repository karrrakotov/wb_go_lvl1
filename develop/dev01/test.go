package main

import "fmt"

/*

=== Задача ===

Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

*/

type Human struct {
	Name string
	Age  int
	Sex  rune
}

func (h *Human) info() {
	fmt.Println("Я - человек.")
}

type Action struct {
	Human
}

func (a *Action) say() {
	fmt.Printf("Меня зовут: %v и я умею говорить.\n", a.Name)
}

func main() {
	human := Human{"Ruslan", 22, 'm'}
	human.info() // Я - человек

	action := Action{human}
	action.info() // Я - человек
	action.say()  // Расширение первоначального метода
}
