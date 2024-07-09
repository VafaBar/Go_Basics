package main

import "fmt"

func main() {
	fmt.Println(many(1, 3))

}
func many(v int, f int) (int, int) {

	return v + f, v - f

}

// 4. Использование нескольких значений return: Создайте функцию, которая
// принимает два числа и возвращает их сумму и разность как два
// отдельных значения.
