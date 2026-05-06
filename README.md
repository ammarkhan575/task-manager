# 📋 Task Manager

A powerful, lightweight command-line task management tool built with **Go**. Manage your tasks efficiently with support for priorities, tags, and status tracking.

---

## ✨ Features

- 🚀 **Fast & Lightweight** - Written in Go for optimal performance
- 📝 **Task Management** - Create, read, update, and delete tasks
- 🎯 **Priority Levels** - Organize tasks by Low, Medium, and High priority
- 🏷️ **Tags** - Categorize tasks with custom tags
- 📊 **Status Tracking** - Track task status (todo, in_progress, done)
- 💾 **Persistent Storage** - Tasks are saved to JSON for data persistence
- 🔍 **Filtering** - Filter tasks by status

---

## 📦 Installation

### Prerequisites
- Go 1.16 or higher

### Clone & Build

```bash
git clone https://github.com/ammarkhan575/task-manager.git
cd task-manager
go build -o task
```

This will create a `task` executable in your current directory.

### Optional: Add to PATH

To use the task manager from anywhere:

```bash
# Copy the binary to a directory in your PATH (e.g., /usr/local/bin)
sudo mv task /usr/local/bin/
```

---

## 🚀 Quick Start

```bash
# Add a new task
./task add -title "Buy groceries" -priority high -desc "Get milk and bread" -tags "shopping,urgent"

# List all tasks
./task list

# Mark a task as done
./task done -id 1

# Delete a task
./task delete -id 1

# Get a specific task
./task get -id 1
```

---

## 📖 Commands

### 1. **Add a Task**

Create a new task with custom properties.

```bash
./task add -title "Task Title" [options]
```

#### Options:
| Option | Type | Required | Description |
|--------|------|----------|-------------|
| `-title` | string | ✅ Yes | The title of the task |
| `-desc` | string | ❌ No | Task description |
| `-priority` | string | ❌ No | Priority level: `low` (default), `medium`, `high` |
| `-tags` | string | ❌ No | Comma-separated tags (e.g., `work,urgent`) |

#### Examples:

```bash
# Simple task
./task add -title "Call client"

# Task with description and priority
./task add -title "Review PR" -desc "Check code quality" -priority high

# Task with all options
./task add -title "Project deadline" -desc "Complete Q1 project" -priority high -tags "work,deadline"

# Task with multiple tags
./task add -title "Meeting prep" -tags "meeting,important,urgent"
```

---

### 2. **List Tasks**

Display all tasks with optional filtering.

```bash
./task list [options]
```

#### Options:
| Option | Type | Description |
|--------|------|-------------|
| `-status` | string | Filter by status: `todo`, `in_progress`, `done` |

#### Examples:

```bash
# Show all tasks
./task list

# Show only pending tasks
./task list -status todo

# Show completed tasks
./task list -status done

# Show tasks in progress
./task list -status in_progress
```

#### Output Format:
```
[ID] | Title | Priority | Status | due: Date
[1] | Buy groceries | High | todo | due: no due date
[2] | Call client | Low | done | due: no due date
```

---

### 3. **Get a Specific Task**

Retrieve detailed information about a specific task.

```bash
./task get -id <task_id>
```

#### Options:
| Option | Type | Required | Description |
|--------|------|----------|-------------|
| `-id` | int | ✅ Yes | The ID of the task to retrieve |

#### Example:

```bash
./task get -id 1

# Output:
# [1] | Buy groceries | High | todo | due: no due date
```

---

### 4. **Mark Task as Done**

Mark a task as complete.

```bash
./task done -id <task_id>
```

#### Options:
| Option | Type | Required | Description |
|--------|------|----------|-------------|
| `-id` | int | ✅ Yes | The ID of the task to mark as done |

#### Example:

```bash
./task done -id 1

# Output:
# ✓ Task #1 marked as done
```

---

### 5. **Delete a Task**

Remove a task permanently.

```bash
./task delete -id <task_id>
```

#### Options:
| Option | Type | Required | Description |
|--------|------|----------|-------------|
| `-id` | int | ✅ Yes | The ID of the task to delete |

#### Example:

```bash
./task delete -id 1

# Output:
# ✓ Task #1 deleted
```

---

## 📊 Task Properties

Each task contains the following information:

```json
{
  "id": 1,
  "title": "Buy groceries",
  "description": "Get milk and bread",
  "priority": 2,
  "status": "todo",
  "created_at": "2024-01-15T10:30:00Z",
  "due_date": null,
  "tags": ["shopping", "urgent"]
}
```

### Property Details:

| Property | Type | Description |
|----------|------|-------------|
| `id` | int | Unique task identifier |
| `title` | string | Task name |
| `description` | string | Detailed description (optional) |
| `priority` | int | 0=Low, 1=Medium, 2=High |
| `status` | string | `todo`, `in_progress`, or `done` |
| `created_at` | timestamp | Task creation time |
| `due_date` | timestamp | Optional deadline |
| `tags` | array | List of tags for categorization |

---

## 💾 Data Storage

Tasks are automatically saved to `tasks.json` in the current directory. The file is created on first use and updated after every operation.

### File Location:
```
./tasks.json
```

### Format:
```json
[
  {
    "id": 1,
    "title": "Sample Task",
    "description": "A sample task",
    "priority": 1,
    "status": "todo",
    "created_at": "2024-01-15T10:30:00Z",
    "tags": ["sample"]
  }
]
```

---

## 🔧 Project Structure

```
task-manager/
├── main.go                 # Entry point
├── go.mod                  # Go module definition
├── README.md              # This file
├── tasks.json             # Task storage (auto-generated)
└── internal/
    ├── cli/
    │   └── cli.go         # CLI command handling
    └── task/
        ├── task.go        # Task model and logic
        ├── store.go       # Task persistence layer
        ├── persistance.go # File I/O operations
        └── errors.go      # Custom error definitions
```

---

## 📝 Usage Workflow

Here's a typical workflow for using the task manager:

```bash
# 1. Create tasks for your project
./task add -title "Design database schema" -priority high -tags "backend,database"
./task add -title "Setup CI/CD pipeline" -priority medium -tags "devops"
./task add -title "Write API documentation" -priority medium -tags "docs"

# 2. View all pending tasks
./task list -status todo

# 3. Start working on the first task
./task done -id 1

# 4. Check completed tasks
./task list -status done

# 5. Filter by another status
./task list -status in_progress

# 6. Get details about a specific task
./task get -id 2

# 7. Remove completed tasks
./task delete -id 1
```

---

## 🎯 Priority Levels

| Level | Value | Description | Use Case |
|-------|-------|-------------|----------|
| **Low** | 0 | Non-urgent tasks | Backlog items, nice-to-haves |
| **Medium** | 1 | Standard priority | Regular work items |
| **High** | 2 | Urgent tasks | Deadlines, blockers |

---

## 🏷️ Tags Best Practices

Use tags to organize tasks by:
- **Project**: `backend`, `frontend`, `devops`
- **Category**: `bug`, `feature`, `documentation`, `testing`
- **Team**: `team-a`, `team-b`, `backend-team`
- **Context**: `urgent`, `review`, `blocked`

### Example:
```bash
./task add -title "Fix login bug" -priority high -tags "bug,urgent,frontend"
```

---

## 🐛 Troubleshooting

### Issue: Command not found
**Solution:** Make sure the task binary is in your PATH or run with `./task`

### Issue: tasks.json not found error
**Solution:** This file is created automatically on first use. Ensure you have write permissions in the current directory.

### Issue: Task not found
**Solution:** Verify the task ID exists by running `./task list`

---

## 🔒 Note

- Each task is automatically assigned a unique ID
- Tasks are saved to `tasks.json` automatically after every operation
- Deleting a task is permanent
- The data file should be backed up for important tasks

---

## 📄 License

This project is open source and available under the MIT License.

---

## 👨‍💻 Author

Created by **Ammar Khan**

---

## 🤝 Contributing

Contributions are welcome! Feel free to submit issues and pull requests.

---

## ⭐ Support

If you find this project helpful, please consider giving it a star!
