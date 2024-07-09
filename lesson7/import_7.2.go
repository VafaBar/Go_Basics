package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	f := time.Now().Second()
	a := math.Sqrt(float64(f))

	fmt.Println(a)
}

// Использование нескольких пакетов: Импортируйте пакеты "math" и
// "time", используйте "math" для вычисления квадратного корня текущего
// времени в секундах.
