package main

import "fmt"

func main() {

	planets := []string{"Earth", "Venus", "Saturn", "Mercury"}

	planets = append(planets, "Pluto", "Jupiter")

	fmt.Println(planets)

	//Создайте слайс, добавьте в него несколько элементов с помощью функции append и выведите все элементы на экран.
}
