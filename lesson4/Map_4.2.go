package main

import "fmt"

func main() {

	regions := map[string]int{
		"Saint-Petersburg": 78,
		"Moscow":           79,
		"Altai Republic":   04,
	}
	regions["Moscow"] = 99

	fmt.Println(regions)

	// Создайте map с несколькими парами ключ-значение.
	// Измените одно из значений и выведите всю map на экран.
}
