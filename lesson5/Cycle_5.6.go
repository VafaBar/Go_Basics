package main

import "fmt"

func main() {

	ages := make(map[string]int)
	ages["Harry"] = 16
	ages["Ron"] = 17
	ages["Hermione"] = 18

	for name, age := range ages {
		fmt.Println(name, age)
	}

}

// Цикл for-range с map: Создайте map с несколькими парами ключзначение.
// Напишите цикл for-range, который выводит каждую пару ключзначение на экран.
