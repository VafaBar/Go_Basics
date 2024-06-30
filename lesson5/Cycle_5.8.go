package main

import "fmt"

func main() {

	n := "I love my parents"
	for _, f := range n {
		d := string(f)
		fmt.Println(d)
	}

}

// Цикл for-range с строкой: Создайте строку. Напишите цикл for-range,
// который выводит каждый символ строки на экран.
