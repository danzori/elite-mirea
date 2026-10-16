package main

import (
	"fmt"
)

func main() {
	var minutes uint

	fmt.Print("Введите количество минут: ")
	_, err := fmt.Scan(&minutes)
	if err != nil {
		fmt.Println("Ошибка: Необходимо ввести целое неотрицательное число")

		return
	}

	hours := minutes / 60
	remainingMinutes := minutes % 60

	fmt.Printf("Полных часов: %d\n", hours)
	fmt.Printf("Оставшихся минут: %d", remainingMinutes)
}
