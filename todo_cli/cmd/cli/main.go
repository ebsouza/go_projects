package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	todo "github.com/ebsouza/todo_app_cli"
)

var todoFileName = ".todo.json"

func listPrinter(l *todo.List, avoidComplete bool, verbose bool) {
	for index, item := range *l {
		if avoidComplete && item.Done {
			continue
		}

		suffix := "[Done]"

		var message string
		if verbose {
			message = fmt.Sprintf("%d - Name: %s - Created at: %s", index+1, item.Task, item.CreatedAt.Format("2024-12-06 15:04:05"))
		} else {
			message = fmt.Sprintf("%d - %s", index+1, item.Task)
		}

		if item.Done {
			fmt.Println(message, suffix)
			continue
		}
		fmt.Println(message)
	}
}

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}

	if len(s.Text()) == 0 {
		return "", fmt.Errorf("task cannot be blank")
	}

	return s.Text(), nil
}

func main() {
	add := flag.Bool("add", false, "Add task to the ToDo List")
	list := flag.Bool("list", false, "List all incomplete tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	delete := flag.Int("delete", 0, "Item to be removed")
	verbose := flag.Bool("verbose", false, "Show complete information")
	avoidComplete := flag.Bool("avoid_complete", false, "Show only incomplete tasks")

	flag.Usage = func() {
		fmt.Fprint(flag.CommandLine.Output(), "Title of help message \n")
		flag.PrintDefaults()
		fmt.Fprint(flag.CommandLine.Output(), "\n >> App developed by ebsouza << \n")
	}

	flag.Parse()

	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		listPrinter(l, *avoidComplete, *verbose)

	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *delete > 0:
		if err := l.Delete(*delete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *add:
		task, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		l.Add(task)

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}
