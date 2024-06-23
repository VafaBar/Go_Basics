package main

import "fmt"

func main() {

	x := map[string]map[int]string{

		"Millonnaya":   {34: "Esenin"},
		"Italyanskaya": {56: "Pushkin"},
		"Rastrelly":    {89: "Lermontov"},
	}

	fmt.Println(x)

	// Создайте map, в которой ключом будет string, а значением map. Map-значение в себе имеет ключ int, значение string. В первой map (с ключом string)
	// ключом будет название улицы, во второй map (которая является значением первой) - ключом будет номер дома,
	// а значением фамилия семьи живущей в нём. Создайте несколько пар ключзначение. Выведите результат.
}
