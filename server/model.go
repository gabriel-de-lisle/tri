package main

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Task struct {
	Text     string
	Priority int32
	Date     time.Time
	Done     bool
}

func (task *Task) SetPriority(p int32) {
	switch p {
	case 1:
		task.Priority = 1
	case 3:
		task.Priority = 3
	default:
		task.Priority = 2
	}
}

func (task *Task) SetDate() {
	task.Date = time.Now()
}

type ByPriority []Task

func (s ByPriority) Len() int          { return len(s) }
func (s ByPriority) Swap(i int, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPriority) Less(i int, j int) bool {
	if s[i].Done != s[j].Done {
		return s[j].Done
	}
	if !s[i].Done {
		if s[i].Priority != s[j].Priority {
			return s[i].Priority < s[j].Priority
		}
		return s[i].Date.After(s[j].Date)
	} else {
		return s[i].Date.After(s[j].Date)
	}

}

func SaveTasks(filename string, tasks []Task) error {
	b, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadTasks(filename string) ([]Task, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Task{}, err
	}

	var tasks []Task
	err = json.Unmarshal(b, &tasks)
	if err != nil {
		return []Task{}, err
	}

	return tasks, nil
}
