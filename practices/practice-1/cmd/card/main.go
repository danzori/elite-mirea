package main

import "fmt"

func main() {
	const course = "Бекенд-разработка на языке Go"
	var name = "Danya"
	group := "Group 1"
	completed := 1
	interest := true

	fmt.Printf("Участник: %s\n", name)
	fmt.Printf("Группа: %s\n", group)
	fmt.Printf("Курс: %s\n", course)
	fmt.Printf("Завершено работ: %d\n", completed)
	fmt.Println("Интерес к бекенду:", interest)

	name = "Sanya"
	completed = 2

	fmt.Printf("Измененный участник: %s\n", name)
	fmt.Printf("Измененное кол-во завершенных работ: %d\n", completed)
}
