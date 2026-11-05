package main

import (
	"fmt"
	"slices"
)

func statusLabel(status string) string {
	switch status {
	case "new":
		return "Новая"
	case "in_progress":
		return "В работе"
	case "done":
		return "Завершена"
	default:
		return "Неизвестный статус"
	}
}

func countStatus(statuses []string, target string) int {
	var count int

	for _, status := range statuses {
		if status == target {
			count++
		}
	}

	return count
}

func findStatus(statuses []string, target string) int {
	return slices.Index(statuses, target)
}

func completionPercent(statuses []string) float64 {
	var count int

	for _, status := range statuses {
		if status == "done" {
			count++
		}
	}

	if count == 0 {
		return 0
	}

	return 100 * float64(count) / float64(len(statuses))
}

func main() {
	statuses := []string{"new", "in_progress", "done", "done"}
	statuses = append(statuses, "new")

	fmt.Printf("Задач всего: %d\n", len(statuses))

	for i, status := range statuses {
		fmt.Printf("%d %s\n", i+1, statusLabel(status))
	}

	fmt.Printf("Завершено: %d\n", countStatus(statuses, "done"))

	counts := make(map[string]int)
	for _, status := range statuses {
		counts[status]++
	}

	fmt.Printf("Статистика: new=%d, in_progress=%d, done=%d\n", counts["new"], counts["in_progress"], counts["done"])

	fmt.Printf("Индекс присутствующего статуса done: %d\n", findStatus(statuses, "done"))
	fmt.Printf("Индекс несуществующего статуса unknown: %d\n", findStatus(statuses, "unknown"))

	fmt.Printf("Доля выполнения задач: %f\n", completionPercent(statuses))
}
