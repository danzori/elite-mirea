package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Task struct {
	ID       int
	Title    string
	Status   string
	Priority int
}

type TaskStore interface {
	Add(title string, priority int) (Task, error)
	Get(id int) (Task, error)
	Complete(id int) error
	Delete(id int) error
	List() []Task
}

var (
	ErrNotFound = errors.New("task not found")
	ErrInvalid  = errors.New("invalid task")
)

var _ TaskStore = (*MemoryStore)(nil)

type MemoryStore struct {
	tasks  []Task
	nextID int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		nextID: 1,
	}
}

func (s *MemoryStore) Add(title string, priority int) (Task, error) {
	title = strings.TrimSpace(title)

	if title == "" || priority < 1 || priority > 3 {
		return Task{}, fmt.Errorf("add task: %w", ErrInvalid)
	}

	task := Task{
		ID:       s.nextID,
		Title:    title,
		Status:   "new",
		Priority: priority,
	}

	s.tasks = append(s.tasks, task)
	s.nextID++

	return task, nil
}

func (s *MemoryStore) Get(id int) (Task, error) {
	for _, task := range s.tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return Task{}, fmt.Errorf("get task %d: %w", id, ErrNotFound)
}

func (s *MemoryStore) Complete(id int) error {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks[i].Status = "done"

			return nil
		}
	}

	return fmt.Errorf("complete task %d: %w", id, ErrNotFound)
}

func (s *MemoryStore) Delete(id int) error {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = slices.Delete(s.tasks, i, i+1)

			return nil
		}
	}

	return fmt.Errorf("delete task %d: %w", id, ErrNotFound)
}

func (s *MemoryStore) List() []Task {
	return slices.Clone(s.tasks)
}

func readLine(scanner *bufio.Scanner, prompt string) (string, bool) {
	fmt.Print(prompt)

	if !scanner.Scan() {
		return "", false
	}

	return strings.TrimSpace(scanner.Text()), true
}

func parseInt(raw string) (int, bool) {
	result, err := strconv.Atoi(raw)

	return result, err == nil
}

func printError(err error) {
	switch {
	case errors.Is(err, ErrNotFound):
		fmt.Println("Такой задачи нет")

	case errors.Is(err, ErrInvalid):
		fmt.Println("Неверные данные")

	default:
		fmt.Printf("Ошибка: %v\n", err)
	}
}

func run(store TaskStore, scanner *bufio.Scanner) {
	defer fmt.Println("Работа приложения завершена")

	for {
		fmt.Println("\n1 — добавить | 2 — список | 3 — завершить | 4 — удалить | 5 - найти | 0 — выход")

		cmd, ok := readLine(scanner, "> ")
		if !ok || cmd == "0" {
			return
		}

		switch cmd {
		case "1":
			title, ok := readLine(scanner, "Название: ")
			if !ok {
				return
			}

			raw, ok := readLine(scanner, "Приоритет (1–3): ")
			if !ok {
				return
			}

			priority, ok := parseInt(raw)
			if !ok {
				fmt.Println("Введите число")

				continue
			}

			task, err := store.Add(title, priority)
			if err != nil {
				printError(err)

				continue
			}

			fmt.Printf("Создана задача #%d\n", task.ID)

		case "2":
			for _, task := range store.List() {
				fmt.Printf("#%d %s [%s], приоритет %d\n", task.ID, task.Title, task.Status, task.Priority)
			}

		case "3", "4":
			raw, ok := readLine(scanner, "ID: ")
			if !ok {
				return
			}

			id, ok := parseInt(raw)
			if !ok {
				fmt.Println("Введите число")

				continue
			}

			var err error
			if cmd == "3" {
				err = store.Complete(id)
			} else {
				err = store.Delete(id)
			}

			if err != nil {
				printError(err)

				continue
			}

			fmt.Println("Готово")

		case "5":
			raw, ok := readLine(scanner, "ID: ")
			if !ok {
				return
			}

			id, err := strconv.Atoi(raw)
			if err != nil {
				fmt.Println("Введите число")

				continue
			}

			task, err := store.Get(id)
			if err != nil {
				printError(err)

				continue
			}

			fmt.Printf("#%d %s [%s], приоритет %d\n", task.ID, task.Title, task.Status, task.Priority)

		default:
			fmt.Println("Неизвестная команда")
		}
	}
}

func main() {
	run(NewMemoryStore(), bufio.NewScanner(os.Stdin))
}
