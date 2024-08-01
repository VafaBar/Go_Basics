package main

import "fmt"

func main() {
	age := new(int)
	fmt.Println(*age)
	*age = 18
	fmt.Println(*age)

	gender := new(string)
	fmt.Println(*gender)
	*gender = "female"
	fmt.Println(*gender)

	isMarried := new(bool)
	*isMarried = true
	fmt.Println(*isMarried)

}

// Инициализация указателей через new (⭐ сложное): Исследуйте, как
// используется функция new для инициализации указателей. Создайте
// несколько переменных с помощью new и измените их значения.
