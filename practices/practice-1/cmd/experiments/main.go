package main

import "fmt"

func main() {
	const language = "Go"
	var name string = "Ирина"
	completed := 1
	var ready bool = true
	var hours float64 = 1.5

	fmt.Println("Изучаем", language)
	fmt.Printf("Участник: %s\n", name)
	fmt.Printf("Выполнено: %d\n", completed)
	fmt.Printf("Готов продолжать: %t\n", ready)
	fmt.Printf("Продолжительность: %.1f часа\n", hours)

	completed = 2
	fmt.Printf("После изменения: %d\n", completed)
}
