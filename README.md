# CLI Todo App

This project is a simple command-line to-do application built with Go. It helps you manage small task lists directly from the terminal without needing a database or external service.

## What it does

- Create a new task from the CLI
- View all saved tasks
- Update a task to mark it as completed
- Store task data in a local JSON file called `task.json`

## How it works

The app presents a menu in the terminal with options to create, list, and update tasks. Each task includes an ID, title, timestamp, and completion status. The list is persisted locally so your tasks remain available across runs.

## Run the app

```bash
go run .
```

This is a lightweight personal productivity tool for tracking tasks in a terminal-based workflow.
