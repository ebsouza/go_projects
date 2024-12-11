package todo

import (
	"os"
	"testing"
)

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
