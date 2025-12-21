package teamdb

import (
	"database/sql"
	"os"
	"testing"
)

func TestTeamDatabase(t *testing.T) {
	const dbFile = "../../data/team.db"
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
	query := "SELECT name FROM sqlite_master WHERE type='table' AND name='teams';"
	err = db.QueryRow(query).Scan(&table_name)
	if err != nil {
		t.Fatalf("Failed to query database schema: %v", err)
	}

	if err == sql.ErrNoRows {
		t.Fatal("Database file was created, but 'teams' table was not found.")
	}
	if err != nil {
		t.Fatalf("Error while querying for 'teams' table: %v", err)
	}
	if table_name != "teams" {
		t.Errorf("Found table, but name was incorrect. Got '%s', expected 'teams'", table_name)
	}

	// ------- Test CreateTeam Function -------
	CreateTeam("The Wonders")
	CreateTeam("Steeves")

	// ------- Test GetTeamName and AssignTeamName
	name, _ := GetTeamName(1)
	t.Logf("Pass: created and aquired team: '%v'", name)
	AssignTeamName(1, "Balling")
	name, _ = GetTeamName(1)
	t.Logf("Pass: changed team name (previously 'The Wonders'), New: '%v'", name)

	// ------- Test GetTeam Funciton -------
	team1 := GetTeam(1)
	t.Logf("Pass: got details of team 1: '%v'", team1)

	// ------- Test GetAllTeams Function -------
	teamids, _ := GetAllTeamIDS()
	t.Logf("Pass: got all team IDS: '%v'", teamids)

	// ------- Test DeleteTeam Function -------
	DeleteTeam(1)
	teamids, _ = GetAllTeamIDS()
	t.Logf("Pass: Deleted team 1. New list of Team IDS: '%v'", teamids)
}
