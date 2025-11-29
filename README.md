# Todo App

A simple API-based todo application built with Go that organizes tasks into multiple categories.

## Features

- RESTful API for managing todos
- Organize todos into multiple categories
- Mark todos as done or incomplete
- Database persistence with SQL
- Docker support for containerized deployment

## Getting Started

### Prerequisites

- Go 1.18 or higher
- Docker and Docker Compose
- Task (installed via scoop)

### Installation

#### 1. Install Scoop (Windows)

If you don't have Scoop installed, follow these steps:

```bash
# Open PowerShell and run:
iwr -useb get.scoop.sh | iex
```

For more details, visit [Scoop's installation guide](https://scoop.sh/)

#### 2. Install Task via Scoop

Once Scoop is installed, you can install Task:

```bash
scoop install task
```

For other operating systems, visit [Task's installation guide](https://taskfile.dev/installation/)

#### 3. Install Docker and Docker Compose

To run Docker and Docker Compose, ensure Docker Desktop is installed on your system. You can download it from [Docker's official website](https://www.docker.com/)

### Running the Application

Once you have Task installed, you can use the Taskfile to run common commands:

```bash
task dev      # Run the application in development mode
task build    # Build the application
task docker   # Build and run using Docker Compose
```

For more available tasks, run:

```bash
task --list
```

## Development

The project uses Task (a task runner) to manage common development tasks. All available tasks are defined in `Taskfile.yml`.

To see all available commands:

```bash
task --list
```

## License

MIT
