package main

import "fmt"

func main() {
	var names1 [4]string

	names1[0] = "Tom"
	names1[1] = "Edgar"
	names1[2] = "Steve"
	names1[3] = "Elvis"

	var names2 [4]string

	names2[0] = "Megan"
	names2[1] = "Agata"
	names2[2] = "Angelina"
	names2[3] = "Brad"

	celebrities := map[string]string{

		names1[0]: "Ford",
		names1[1]: "Poe",
		names1[2]: "Jobs",
		names1[3]: "Presley",
		names2[0]: "Fox",
		names2[1]: "Christie",
		names2[2]: "Joli",
		names2[3]: "Pitt",
	}
	fmt.Println(celebrities)

	// Задание: Создайте два массива одного типа с одинаковой длиной. Создайте map, в которой ключом являются созданные массивы, тип значений задайте произвольно,
	// попробуйте наполнить созданную map разными данными. Попробуйте использовать слайс как ключ в мапе, объясните итог.
	// Ответ: Слайсы нельзя использовать в качестве ключей в мапе, так как слайсы являются несравнимым типом
	// (этот тип не сравнить с помощью операторов == или !=).
}
