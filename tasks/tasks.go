package task

import "fmt"

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

func ListTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks added.")
		return
	}

	for _, task := range tasks {
		status := " "
		if task.Complete {
			status = " "
			if task.Complete {
				status = "✓"
			}
			fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Name)
		}
	}
}
