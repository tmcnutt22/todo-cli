package actions

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func GetTodos(filename string) (*Todos, error) {
	// open a filepath from the os
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	todos := &Todos{}

	for _, record := range records {
		if len(record) < 2 {
			return nil, fmt.Errorf("Invalid")
		}

		dateTime, err := time.Parse("2006-01-02 15:04:05", record[0])
		if err != nil {
			fmt.Println("could not parse date")
			return nil, err
		}

		todo := Todo{
			DateTime: dateTime,
			Message:  record[1],
		}

		todos.Items = append(todos.Items, todo)
	}

	return todos, nil
}
