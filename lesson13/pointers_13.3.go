package main

import "fmt"

func Pointer(x int) *int {
	p := new(int)
	*p = x
	return p
}

func main() {

	p1 := Pointer(8)
	fmt.Println("p1:", *p1)

}

// Возвращение указателя из функции: Напишите функцию, которая
// возвращает указатель на новую переменную int.
