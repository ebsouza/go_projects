package todo

import (
	"fmt"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) Add(task string) {
	i := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, i)
}

func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d doesn't exist", i)
	}

	itemIndex := i - 1
	ls[itemIndex].Done = true
	ls[itemIndex].CompletedAt = time.Now()

	return nil
}

func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d doesn't exist", i)
	}

	itemIndex := i - 1
	nextItemIndex := i
	*l = append(ls[:itemIndex], ls[nextItemIndex:]...)

	return nil
}
