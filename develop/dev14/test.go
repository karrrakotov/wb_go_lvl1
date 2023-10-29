package main

/*

=== Задача ===

Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

*/

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt interface{} = 5
	var myString interface{} = "Hello"
	var myBool interface{} = true
	var myChannel interface{} = make(chan int)

	// Проверка типа с помощью type assertion
	fmt.Println("Type of myInt:", getType(myInt))
	fmt.Println("Type of myString:", getType(myString))

	// Проверка типа с помощью reflect
	fmt.Println("Type of myBool:", getTypeWithReflect(myBool))
	fmt.Println("Type of myChannel:", getTypeWithReflect(myChannel))
}

// 1 Способ, использование - type assertion
func getType(data interface{}) string {
	switch data.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	default:
		return "unknown type"
	}
}

// 1 Способ, использование - reflect
func getTypeWithReflect(data interface{}) string {
	switch reflect.ValueOf(data).Kind() {
	case reflect.Int:
		return "int"
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Chan:
		return reflect.Chan.String()
	default:
		return "unknown type"
	}
}
