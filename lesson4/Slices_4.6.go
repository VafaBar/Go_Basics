package main

import "fmt"

func main() {

	planets := []string{"Earth", "Venus", "Saturn", "Mercury"}

	sub_planets := planets[1:3]

	fmt.Println(sub_planets)

	//  Создайте слайс с несколькими элементами, затем создайте новый слайс как срез первого слайса и выведите его на экран.
}
