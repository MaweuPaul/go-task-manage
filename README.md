# Task Manager API

A simple REST API built with Go and Gin framework — my first Go project.

## Tech Stack
- Go
- Gin
- UUID generation

## Getting Started

### Prerequisites
- Go 1.21+

### Installation
```bash
git clone git@github.com:MaweuPaul/go-task-manage.git
cd taskmanager
go mod tidy
go run main.go
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | /tasks | Create a task |
| GET | /tasks | Get all tasks |
| GET | /tasks/:id | Get a single task |
| PUT | /tasks/:id | Update a task |
| DELETE | /tasks/:id | Delete a task |

## Example Request

```json
{
    "title": "Buy milk",
    "description": "From the shop",
    "status": "pending"
}
```


