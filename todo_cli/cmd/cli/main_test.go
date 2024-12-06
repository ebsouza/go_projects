package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool ...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests ...")

	result := m.Run()

	fmt.Println("Cleaning up ...")
	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	task1 := "Task 1"
	task2 := "Task 2"
	t.Run("AddTaskFromArguments", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", task1)

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}

		cmd = exec.Command(cmdPath, "-add", task2)

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	task3 := "Task 3"
	t.Run("AddTaskFromSTDIN", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add")
		cmdStdIn, err := cmd.StdinPipe()
		if err != nil {
			t.Fatal(err)
		}
		io.WriteString(cmdStdIn, task3)
		cmdStdIn.Close()

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("CompleteTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-complete", "3")

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("DeleteTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-delete", "1")

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatal(err)
		}

		expected := fmt.Sprintf("1 - %s\n2 - %s [X]\n", task2, task3)

		if expected != string(out) {
			t.Errorf("Expected %q, got %q instead\n", expected, string(out))
		}
	})

	t.Run("ListTasksVerbose", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list", "-verbose")
		_, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasksAvoidComplete", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list", "-avoid_complete")
		out, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatal(err)
		}

		expected := fmt.Sprintf("1 - %s\n", task2)

		if expected != string(out) {
			t.Errorf("Expected %q, got %q instead\n", expected, string(out))
		}
	})

}
