package main

import (
	"fmt"
	mr "math/rand"
)

func main() {
	fmt.Println(mr.Intn(1000))
}

// Использование псевдонимов для пакетов: Импортируйте пакет
// "math/rand" с псевдонимом и используйте его для генерации случайного
// числа.
