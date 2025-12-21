package taskdb

import (
	"database/sql"
	"os"
	"testing"
)

func TestTaskDatabase(t *testing.T) {
	const dbFile = "../../data/task.db"
	os.Remove(dbFile)

	// t.Cleanup(func() {
	// 	err := os.Remove(dbFile)
	// 	if err != nil {
	// 		t.Logf("Warning: failed to remove previous database: %v", err)
	// 	}
	// })

	// ------- Testing InitDatabase Function -------

	InitDatabase()

	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		t.Errorf("Database file was not created")
	}

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		t.Fatalf("Failed to open database file: %v", err)
	}
	defer db.Close()

	var table_name string
	query := "SELECT name FROM sqlite_master WHERE type='table' AND name='tasks';"
	err = db.QueryRow(query).Scan(&table_name)
	if err != nil {
		t.Fatalf("Failed to query database schema: %v", err)
	}
	if err == sql.ErrNoRows {
		t.Fatal("Database file was created, but 'tasks' table was not found.")
	}
	if err != nil {
		t.Fatalf("Error while querying for 'tasks' table: %v", err)
	}
	if table_name != "tasks" {
		t.Errorf("Found table, but name was incorrect. Got '%s', expected 'tasks'", table_name)
	}

	CreateTask("Wash Dishes")
	CreateTask("Clean Sink")
	CreateTask("Do Laundry")
	CreateTask("Put Away Dishes")

	AssignTaskHiddenStatus(4, true)
	AssignTaskDescription(1, "Handwash dishes and load dishwasher.")
	AssignTaskPoints(1, 3)
	AssignTaskDuedate(1, 999999993)
	AssignTaskTimeframe(1, "3pm-4pm")
	AssignTaskTitle(1, "Wash and load dishwasher")
	AssignTaskTriggers(1, 4)

	ids, _ := GetAllTaskIDS()

	for _, id := range ids {
		t.Logf("PASS: Got task: '%v'", GetTask(id))
	}

	DeleteTask(2)

	ids, _ = GetAllTaskIDS()

	for _, id := range ids {
		t.Logf("PASS: Got task: '%v'", GetTask(id))
	}
}
