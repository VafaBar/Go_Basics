package main

import "fmt"

func main() {
	myV := 5
	fmt.Println(myV)

	increm(&myV)
	fmt.Println(myV)
}

func increm(myV *int) {
	*myV++
	fmt.Println(*myV)
}

// Передача указателя в функцию: Напишите функцию, которая принимает
// указатель на int и инкрементирует значение, на которое указывает
// указатель.
