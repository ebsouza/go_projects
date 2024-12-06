package todo

import (
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	l := List{}

	taskName := "Task"
	l.Add(taskName)

	if len(l) != 1 {
		t.Errorf("Expected %d, got %d instead.", 1, len(l))
	}

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := List{}

	taskName := "Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("New task should be not completed")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Errorf("New task should be completed")
	}

}

func TestDelete(t *testing.T) {
	l := List{}

	taskNames := []string{
		"Task 1",
		"Task 2",
		"Task 3",
	}

	for _, v := range taskNames {
		l.Add(v)
	}

	for index, v := range taskNames {
		if l[index].Task != v {
			t.Errorf("Expected %q, got %q instead.", v, l[index].Task)
		}
	}

	if len(l) != len(taskNames) {
		t.Errorf("Expected %d, got %d instead.", len(taskNames), len(l))
	}

	l.Delete(2)

	if len(l) != len(taskNames)-1 {
		t.Errorf("Expected %d, got %d instead.", len(taskNames)-1, len(l))
	}

	if l[1].Task != taskNames[2] {
		t.Errorf("Expected %q, got %q instead.", taskNames[2], l[1].Task)
	}

}

func TestSaveGet(t *testing.T) {
	l1 := List{}
	l2 := List{}

	tf, err := os.CreateTemp("", "")

	if err != nil {
		t.Errorf("Error creating temporary file: %s", err)
	}

	defer os.Remove(tf.Name())

	l1.Add("Task 1")
	l1.Add("Task 2")

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving l1 to file: %s", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting l1 from file: %s", err)
	}

	if len(l1) != len(l2) {
		t.Errorf("Expected: equal length; Got: len(l1) = %d  len(l2) = %d", len(l1), len(l2))
	}
}
