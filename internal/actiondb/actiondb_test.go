package actiondb

import (
	"database/sql"
	"os"
	"testing"
)

func TestTaskDatabase(t *testing.T) {
	const dbFile = "../../data/action.db"
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
	query := "SELECT name FROM sqlite_master WHERE type='table' AND name='actions';"
	err = db.QueryRow(query).Scan(&table_name)
	if err != nil {
		t.Fatalf("Failed to query database schema: %v", err)
	}
	if err == sql.ErrNoRows {
		t.Fatal("Database file was created, but 'actions' table was not found.")
	}
	if err != nil {
		t.Fatalf("Error while querying for 'actions' table: %v", err)
	}
	if table_name != "actions" {
		t.Errorf("Found table, but name was incorrect. Got '%s', expected 'actions'", table_name)
	}

	// Testing other functions

	LogAction(1, "Wash Dishes", 3)
	LogAction(3, "Clean Counters", 2)

	timestamps, _ := GetAllTimestamps()

	action, _ := GetAction(timestamps[1])
	t.Logf("Pass: Created actions, got timestamps, and listed action 2: '%v'", action)
}
