package models

import (
	"encoding/json"
	"os"
	"sync"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

var NextID = 1

const dbFile = "students.json"

var mu sync.Mutex

func SaveToFile(students []Student) error {
	mu.Lock()
	defer mu.Unlock()
	file, err := os.OpenFile(dbFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(students)
}

func LoadFromFile() ([]Student, error) {
	mu.Lock()
	defer mu.Unlock()
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return []Student{}, nil
	}
	file, err := os.Open(dbFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var students []Student
	err = json.NewDecoder(file).Decode(&students)
	return students, err
}
