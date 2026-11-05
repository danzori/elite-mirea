package main

import "fmt"

type Task struct {
	ID       int
	Title    string
	Status   string
	Priority int
}

func (t *Task) Summary() string {
	return fmt.Sprintf("#%d %s [%s], приоритет %d", t.ID, t.Title, t.Status, t.Priority)
}

func (t *Task) Complete() {
	t.Status = "done"
}

func (t *Task) SetPriority(priority int) bool {
	if priority < 1 || priority > 3 {
		return false
	}

	t.Priority = priority

	return true
}

func findTaskByID(tasks []Task, id int) *Task {
	for _, task := range tasks {
		if task.ID == id {
			return &task
		}
	}

	return nil
}

func tasksByStatus(tasks []Task, status string) []Task {
	var result []Task

	for _, task := range tasks {
		if task.Status == status {
			result = append(result, task)
		}
	}

	return result
}

func main() {
	tasks := []Task{
		{ID: 1, Title: "Установить Go", Status: "done", Priority: 1},
		{ID: 2, Title: "Освоить функции", Status: "in_progress", Priority: 2},
		{ID: 3, Title: "Изучить структуры", Status: "new", Priority: 3},
	}

	foundTask := findTaskByID(tasks, 2)
	if foundTask != nil {
		foundTask.Complete()
		foundTask.SetPriority(3)
	}

	var completedCount int

	for _, task := range tasks {
		if task.Status == "done" {
			completedCount++
		}

		fmt.Println(task.Summary())
	}

	fmt.Printf("Задач завершено: %d", completedCount)

	if findTaskByID(tasks, 99) == nil {
		fmt.Println("Задача не найдена")
	}
}
