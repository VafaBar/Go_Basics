package main

import "fmt"

func main() {
	var sl []*int = make([]*int, 3)
	x, y, z := 10, 20, 30
	sl[0] = &x
	sl[1] = &y
	sl[2] = &z
	sl = []*int{&x, &y, &z}
	fmt.Println(*sl[0], *sl[1], *sl[2])

	for _, slice := range sl {
		*slice = *slice * 2

	}
	fmt.Println(*sl[0], *sl[1], *sl[2])
}

// func Double (sl int){
// 	for_, c:= range sl

// fmt.Println (Double)
//  }

// Указатели и циклы: Создайте слайс из указателей на числа. Используйте
// цикл for-range для изменения каждого числа, на которое указывает
// указатель, каждое значение увеличьте в два раза.
