package actions

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type Todos struct {
	Items []Todo
}

type Todo struct {
	DateTime time.Time
	Message  string
}

func NewTodo(message string) *Todo {
	t := Todo{DateTime: time.Now(), Message: message}

	return &t
}

func writeCSV(filepath string, todos []Todo) error {
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, todo := range todos {
		// Format date for CSV
		dateStr := todo.DateTime.Format("2006-01-02 15:04:05")
		err := writer.Write([]string{dateStr, todo.Message})
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

func SaveTodos(todos *Todos, newTodo *Todo) {
	todos.Items = append(todos.Items, *newTodo)

	err := writeCSV("~/temp", todos.Items)
	if err != nil {
		fmt.Println("Shit didn't save")
	}
}
