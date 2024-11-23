package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Storage struct {
	filename string
}

func NewStorage(filename string) *Storage {
	return &Storage{
		filename: filename,
	}
}

func (storage *Storage) Load(tasks *Tasks) error {
	if _, err := os.Stat(storage.filename); errors.Is(err, os.ErrNotExist) {
		os.Create(storage.filename)
	}
	jsonData, _ := os.ReadFile(storage.filename)

	return json.Unmarshal(jsonData, &tasks)
}

func (storage *Storage) Save(tasks *Tasks) error {
	jsonData, err := json.MarshalIndent(tasks, "", "    ")

	if err != nil {
		fmt.Println(err)
		return err
	}

	return os.WriteFile(storage.filename, jsonData, 0644)
}
