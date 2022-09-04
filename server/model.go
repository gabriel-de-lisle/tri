package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Task struct {
	gorm.Model
	Description string
	Priority    int8
	Done        bool
}

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&Task{})
}

func CreateTask(description string, priority int) error {
	newTask := Task{
		Description: description,
		Done:        false,
	}

	switch priority {
	case 1:
		newTask.Priority = 1
	case 3:
		newTask.Priority = 3
	default:
		newTask.Priority = 2
	}

	DB.AutoMigrate(&Task{})

	DB.Create(&newTask)

	return nil
}

func MarkTasksAsDone(taskIds []uint) error {
	result := DB.Model(&Task{}).Where("ID IN ?", taskIds).Update("Done", true)

	return result.Error
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
		return s[i].UpdatedAt.After(s[j].UpdatedAt)
	} else {
		return s[i].UpdatedAt.After(s[j].UpdatedAt)
	}

}

func GetTasks(showAll bool, showDone bool) ([]Task, error) {
	var result *gorm.DB
	var tasks []Task

	if showAll {
		result = DB.Find(&tasks)
	} else {
		result = DB.Where("Done = ?", showDone).Find(&tasks)
	}

	return tasks, result.Error
}
