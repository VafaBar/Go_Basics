package main

import "fmt"

func main() {

	spells := []string{"AvadaKedavra", "Crucio", "Imperio", "Aguamenti", "Alohomora", "Bombardo"}

	fmt.Println(cap(spells))

	spells = append(spells, "Sectumsempra", "Bombardomaxima")

	fmt.Println(cap(spells))

	spells = append(spells, "Lumos", "Nox")

	fmt.Println(cap(spells))

	spells = append(spells, "Arania Exumai", "Arresto Momentum", "Vipera Evanesca ")
	fmt.Println(cap(spells))

	spells = append(spells, "Reducio", "Wingardium Leviosa ", "Oculus Reparo", "Obliviate", "Obscuro", "Oppugno", "Periculum")
	fmt.Println(cap(spells))

	fmt.Println(spells)

	// Создайте слайс произвольного типа с 6 элементами, узнайте его вместимость с помощью функции cap, опытным
	// путём определите в каком случае вместимость слайса будет 12, 24 и более.

}
