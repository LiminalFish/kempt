package userdb

import (
	"database/sql"
	"os"
	"testing"
)

func TestInitDatabase(t *testing.T) {
	const dbFile = "user.db"
	err := os.Remove(dbFile)
	if err != nil {
		t.Logf("Warning: failed to remove previous database: %v", err)
	}

	// t.Cleanup(func() {
	// 	err := os.Remove(dbFile)
	// 	if err != nil {
	// 		t.Logf("Warning: failed to remove database: %v", err)
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
	query := "SELECT name FROM sqlite_master WHERE type='table' AND name='users';"
	err = db.QueryRow(query).Scan(&table_name)
	if err != nil {
		t.Fatalf("Failed to query database schema: %v", err)
	}

	if err == sql.ErrNoRows {
		t.Fatal("Database file was created, but 'users' table was not found.")
	}
	if err != nil {
		t.Fatalf("Error while querying for 'users' table: %v", err)
	}
	if table_name != "users" {
		t.Errorf("Found table, but name was incorrect. Got '%s', expected 'users'", table_name)
	}

	// ------- Testing CreateUser Function -------

	CreateUser("Bobby", 1, "hieo00w0w00w0w0")

	var user_id string
	query = "SELECT userid FROM users WHERE name = 'Bobby'"
	err = db.QueryRow(query).Scan(&user_id)
	if err != nil {
		t.Fatalf("Error while selecting user from table: %v", err)
	}
	t.Logf("Pass: found user ID '%v' from table '%v'.", user_id, table_name)

	CreateUser("Nina", 1, "oijwefoijweoij")
	CreateUser("Jorge Sanchez", 1, "MMMMMMMMMM")

	query = "SELECT userid FROM users WHERE name = 'Nina'"
	err = db.QueryRow(query).Scan(&user_id)
	if err != nil {
		t.Fatalf("Error while selecting user from table: '%v'", err)
	}
	if user_id != "2" {
		t.Errorf("Expected ID 2 for second user. Got: '%v'", user_id)
	}
	t.Logf("Pass: found user ID '%v' from table '%v'.", user_id, table_name)

	query = "SELECT userid FROM users WHERE name = 'Jorge Sanchez'"
	err = db.QueryRow(query).Scan(&user_id)
	if err != nil {
		t.Fatalf("Error while selecting user from table: '%v'", err)
	}
	if user_id != "3" {
		t.Errorf("Expected ID 3 for third user. Got: '%v'", user_id)
	}
	t.Logf("Pass: found user ID '%v' from table '%v'.", user_id, table_name)

	// ------- Testing GetTeamid & AssignTeamid Function -------

	teamid, err := GetTeamid(1)
	if teamid != 0 {
		t.Errorf("Expected team 0 (unnasigned). Got: '%v'", teamid)
	}
	AssignTeamid(1, 1)
	teamid, err = GetTeamid(1)
	if teamid != 1 {
		t.Fatalf("Expected team 1. Got: '%v' from user 1", teamid)
	}
	t.Logf("Pass: Team successfully updated. Got: '%v' from user 1", teamid)
	AssignTeamid(3, 1)
	teamid, err = GetTeamid(3)
	if teamid != 1 {
		t.Fatalf("Expected team 1. Got: '%v' from user 3", teamid)
	}
	t.Logf("Pass: Team successfully updated. Got: '%v' from user 3", teamid)
}
