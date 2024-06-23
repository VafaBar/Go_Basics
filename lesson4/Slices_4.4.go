package main

import "fmt"

func main() {

	planets := make([]string, 3, 10)
	fmt.Println(planets)

	planets = append(planets, "Earth", "Venus", "Saturn", "Mercury")

	fmt.Println(planets)

	// Создайте слайс с помощью функции make, добавьте в него несколько элементов и выведите эти элементы на экран.
}
