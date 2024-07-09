package main

import "fmt"

func main() {
	fmt.Println(Dobby(Malfoy))

}

func Dobby(fn func(int) int) int {
	b := fn(20)
	return b
}

func Malfoy(f int) int {
	return f * 2
}

// Функция как аргумент другой функции (сложное): Исследуйте,
// можем ли мы в Go передать функцию как аргумент другой функции. Если
// нет, то объясните, как можно достичь похожего результата.
