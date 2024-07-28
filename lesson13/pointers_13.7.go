package main

import "fmt"

func main() {
	var age *int
	if age == nil {
		fmt.Println("age is nil")
	} else {
		fmt.Println(*age)
	}

}

// Указатели и nil (⭐ сложное): Исследуйте, что произойдет, если
// попытаться создать переменную с типом *int а затем разыменовать ее.
// Придумайте вариант, который безопасно обрабатывает такую ситуацию.
