package main

import "fmt"

func main() {

	subwaystation := []string{"Mayakovskaya", "Dybenko", "Gorkovskaya", "Sennaya"}

	subwaystation[2] = "Sportivnaya"
	fmt.Println(subwaystation)
	// Создайте слайс с несколькими строками, измените одно из значений и выведите все элементы на экран.
}
