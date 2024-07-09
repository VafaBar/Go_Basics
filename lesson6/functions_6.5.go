package main

import "fmt"

func main() {
	change("Robert", "Adolf", "Napoleon")
}
func change(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}

}

// Функция с переменным числом аргументов: Создайте функцию, которая
// принимает переменное количество аргументов и выводит их на экран.
