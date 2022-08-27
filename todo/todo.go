package todo

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Item struct {
	Text     string
	Priority int
	Date     time.Time
	Done     bool
}

func (item *Item) SetPriority(p int) {
	switch p {
	case 1:
		item.Priority = 1
	case 3:
		item.Priority = 3
	default:
		item.Priority = 2
	}
}

func (item *Item) SetDate() {
	item.Date = time.Now()
}

type ByPriority []Item

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

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}

	var items []Item
	err = json.Unmarshal(b, &items)
	if err != nil {
		return []Item{}, err
	}

	return items, nil
}
