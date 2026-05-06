package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ammarkhan575/task-manager/internal/task"
)

const dataFile = "tasks.json"

func Run() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	store := task.NewStore()
	if err := store.Load(dataFile); err != nil {
		fmt.Fprintf(os.Stderr, "error loading tasks: %v\n", err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		runAdd(store)
	case "list":
		runList(store)
	case "get":
		runGet(store)
	case "delete":
		runDelete(store)
	case "done":
		runDone(store)
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}

	// Save after every mutating command
	if err := store.Save(dataFile); err != nil {
		fmt.Fprintf(os.Stderr, "error saving tasks: %v\n", err)
		os.Exit(1)
	}
}

func runAdd(store *task.Store) {
	// Each subcommand gets its own FlagSet - isolated flags
	cmd := flag.NewFlagSet("add", flag.ExitOnError)
	title := cmd.String("title", "", "Title of the task (required)")
	description := cmd.String("desc", "", "Description of the task")
	priority := cmd.String("priority", "low", "low|medium|high")
	tags := cmd.String("tags", "", "Comma-separated tags")

	cmd.Parse(os.Args[2:]) // parse everything after "add"

	if *title == "" {
		fmt.Fprintln(os.Stderr, "error: -title is required")
		cmd.Usage()
		os.Exit(1)
	}

	p := parsePriority(*priority)
	t := task.NewTask(store.NextID(), *title, *description, p)
	if *tags != "" {
		t.AddTags(strings.Split(*tags, ",")...)
	}
	store.AddTask(t)
	fmt.Printf("✓ Added task #%d: %s\n", t.ID, t.Title)

}

func parsePriority(s string) task.Priority {
	switch s {
	case "medium":
		return task.PriorityMedium
	case "high":
		return task.PriorityHigh
	default:
		return task.PriorityLow
	}
}

func runList(store *task.Store) {
	cmd := flag.NewFlagSet("list", flag.ExitOnError)
	statusFilter := cmd.String("status", "", "filter by status: todo|in_progress|done")
	cmd.Parse(os.Args[2:])

	tasks := store.GetAll()
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	for _, t := range tasks {
		if *statusFilter != "" && string(t.Status) != *statusFilter {
			continue
		}
		fmt.Println(t) // calls t.String() via Stringer interface
	}
}

func runDone(store *task.Store) {
	cmd := flag.NewFlagSet("done", flag.ExitOnError)
	id := cmd.Int("id", 0, "task ID to mark complete")
	cmd.Parse(os.Args[2:])

	t, err := store.GetByID(*id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	t.Complete()
	fmt.Printf("✓ Task #%d marked as done\n", t.ID)
}

func runDelete(store *task.Store) {
	cmd := flag.NewFlagSet("delete", flag.ExitOnError)
	id := cmd.Int("id", 0, "task ID to delete")
	cmd.Parse(os.Args[2:])

	if !store.Delete(*id) {
		fmt.Fprintf(os.Stderr, "error: task #%d not found\n", *id)
		os.Exit(1)
	}
}

func runGet(store *task.Store) {
	cmd := flag.NewFlagSet("get", flag.ExitOnError)
	id := cmd.Int("id", 0, "task ID to get")
	cmd.Parse(os.Args[2:])
	
	t, err := store.GetByID(*id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(t)
}

func printUsage() {
	fmt.Println(`Usage: task <command> [flags]

Commands:
  add    -title "..." [-priority low|medium|high] [-desc "..."] [-tags "a,b"]
  list   [-status todo|in_progress|done]
  done   -id <n>
  get    -id <n>
  delete -id <n>`)
}
