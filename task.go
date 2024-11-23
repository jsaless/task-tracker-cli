package main

import (
	"errors"
	"fmt"
	"time"
)

type Task struct {
	Id          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

type Tasks []Task

func (tasks *Tasks) getIndexById(id int) (int, error) {
	for index, task := range *tasks {
		if task.Id == id {
			return index, nil
		}
	}

	err := errors.New("Task ID not found")
	fmt.Println(err)

	return -1, err
}

func (tasks *Tasks) add(description string) {
	taskId := 1
	t := *tasks

	if len(t) > 0 {
		taskId = t[len(t)-1].Id + 1
	}

	task := &Task{
		Id:          taskId,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	}

	*tasks = append(t, *task)
	fmt.Printf("Task added successfully (ID: %d)\n", taskId)
}

func (tasks *Tasks) delete(id int) error {
	t := *tasks

	taskIndex, err := t.getIndexById(id)

	if err != nil {
		return err
	}

	firstSlice := t[:taskIndex]
	secondSlice := t[taskIndex+1:]

	*tasks = append(firstSlice, secondSlice...)

	return nil
}

func (tasks *Tasks) update(id int, description string) error {
	t := *tasks

	taskIndex, err := t.getIndexById(id)

	if err != nil {
		return err
	}

	updatedTime := time.Now()

	t[taskIndex].Description = description
	t[taskIndex].UpdatedAt = &updatedTime

	return nil
}

func (tasks *Tasks) status(id int, status string) error {
	t := *tasks

	taskIndex, err := t.getIndexById(id)

	if err != nil {
		return err
	}

	updatedTime := time.Now()

	t[taskIndex].Status = status
	t[taskIndex].UpdatedAt = &updatedTime

	return nil
}

func (tasks *Tasks) list(filter string) {

	switch filter {
	case "done":
		for _, task := range *tasks {
			if task.Status == "done" {
				fmt.Printf("ID: %d | Description: %s | Status: %s | Created At: %s | Updated At: %s\n", task.Id, task.Description, task.Status, task.CreatedAt.Format(time.ANSIC), task.UpdatedAt.Format(time.ANSIC))
			}
		}
	case "todo":
		for _, task := range *tasks {
			if task.Status == "todo" {
				fmt.Printf("ID: %d | Description: %s | Status: %s | Created At: %s | Updated At: %s\n", task.Id, task.Description, task.Status, task.CreatedAt.Format(time.ANSIC), task.UpdatedAt.Format(time.ANSIC))
			}
		}
	case "in-progress":
		for _, task := range *tasks {
			if task.Status == "in-progress" {
				fmt.Printf("ID: %d | Description: %s | Status: %s | Created At: %s | Updated At: %s\n", task.Id, task.Description, task.Status, task.CreatedAt.Format(time.ANSIC), task.UpdatedAt.Format(time.ANSIC))
			}
		}
	case "all":
		for _, task := range *tasks {
			fmt.Printf("ID: %d | Description: %s | Status: %s | Created At: %s | Updated At: %s\n", task.Id, task.Description, task.Status, task.CreatedAt.Format(time.ANSIC), task.UpdatedAt.Format(time.ANSIC))
		}
	}
}
