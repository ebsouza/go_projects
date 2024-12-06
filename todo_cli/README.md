### TODO CLI App

![](readme/todo-app.png)

#### Initializing the module
```
go mod init github.com/ebsouza/todo_app_cli
```

#### Building the application
```
go build -o todo_app cmd/cli/main.go
```

#### Commands

##### 1. Add task

```
./todo_app -add Task1
```

```
echo "Task1" | ./todo_app -add
```

##### 2. List task
Consider using *-verbose* or *-avoid_complete* flags when listing tasks
```
./todo_app -list
```

##### 3. Mark as complete

```
./todo_app -complete <Task Index>
```

##### 4. Delete a task
```
./todo_app -delete <Task Index>
```

##### 5. Help instructions
```
./todo_app -h
```