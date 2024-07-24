package main

import "strconv"

func main() {

}
func MyFn() (int, error) {
	StrVl := "12"
	IntVl, errorik := strconv.Atoi(StrVl)
	if errorik != nil {
		return 0, errorik
	}
	return IntVl, nil

}

// Обработка ошибки из стандартной библиотеки: Вызовите функцию из
// стандартной библиотеки Go, которая может вернуть ошибку (например,
// os.Open). Обработайте возвращаемую ошибку.
