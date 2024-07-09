package main

import "fmt"

func main() {
	n1()
}

func n1() {
	fmt.Println("Функция №1")
	n2()
}

func n2() {
	fmt.Println("Я вызвана ")

}

// Вызов функции из другой функции: Создайте две функции и вызовите
// одну из них из другой.
