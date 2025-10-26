package userdb

import (
	"database/sql"
	"os"
	"testing"
)

func TestUserDatabase(t *testing.T) {
	const dbFile = "../../data/user.db"
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

	// ------- Testing GetUserName & AssignUserName Functions -------
	name, err := GetUserName(1)
	if name != "Bobby" {
		t.Fatalf("Failed to aquire name of user 1. Expected 'Bobby'. Got: '%v'", name)
	}
	t.Logf("Pass: Successfully aquired name of user 1: '%v'", name)

	AssignUserName(1, "Bobbert")
	name, err = GetUserName(1)
	if name != "Bobbert" {
		t.Fatalf("Failed to change name of user 1. Expected 'Bobbert'. Got: '%v'", name)
	}
	t.Logf("Pass: Successfully changed name of user 1. Previous: Bobby; New: '%v'", name)

	// ------- Testing GetUserType & AssignUserType Functions -------
	usertype, err := GetUserType(1)
	if usertype != 1 {
		t.Fatalf("Failed to aquire administration level of user 1. Expected '1'. Got: '%v'", usertype)
	}
	t.Logf("Pass: Successfully aquired administrator level of user 1: '%v'", usertype)

	AssignUserType(1, 2)
	usertype, err = GetUserType(1)
	if usertype != 2 {
		t.Fatalf("Failed to change administration level of user 1. Expected '2'. Got: '%v'", usertype)
	}
	t.Logf("Pass: Successfully changed administration level of user 1. Previous: 1; New: '%v'", usertype)

	// ------- Testing GetUserPoints & AssignUserPoints Functions -------
	points, err := GetUserPoints(1)
	if points != 0 {
		t.Fatalf("Failed to aquire points of user 1. Expected '0'. Got: '%v'", points)
	}
	t.Logf("Pass: Successfully aquired points of user 1: '%v'", points)

	AssignUserPoints(1, 5)
	points, err = GetUserPoints(1)
	if points != 5 {
		t.Fatalf("Failed to change points of user 1. Expected '5'. Got: '%v'", points)
	}
	t.Logf("Pass: Successfully changed points of user 1. Previous: 0; New: '%v'", points)

	// ------- Testing GetUserRank & AssignUserRank Functions -------
	rank, err := GetUserRank(1)
	if rank != -1 {
		t.Fatalf("Failed to aquire rank of user 1. Expected '-1'. Got: '%v'", rank)
	}
	t.Logf("Pass: Successfully aquired rank of user 1: '%v'", rank)

	AssignUserRank(1, 1)
	rank, err = GetUserRank(1)
	if rank != 1 {
		t.Fatalf("Failed to change rank of user 1. Expected '1'. Got: '%v'", rank)
	}
	t.Logf("Pass: Successfully changed rank of user 1. Previous: -1; New: '%v'", rank)

	// ------- Testing GetUserToken Function -------
	token, err := GetUserToken(1)
	if token != "hieo00w0w00w0w0" {
		t.Fatalf("Failed to aquire token of user 1. Expected 'hieo00w0w00w0w0'. Got: '%v'", token)
	}
	t.Logf("Pass: Successfully aquired token of user 1: '%v'", token)

	// ------- Testing GetTeamid & AssignTeamid Functions -------

	teamid, err := GetUserTeamid(1)
	if teamid != 0 {
		t.Errorf("Expected team 0 (unnasigned). Got: '%v'", teamid)
	}
	AssignUserTeamid(1, 1)
	teamid, err = GetUserTeamid(1)
	if teamid != 1 {
		t.Fatalf("Expected team 1. Got: '%v' from user 1", teamid)
	}
	t.Logf("Pass: Team successfully updated. Got: '%v' from user 1", teamid)
	AssignUserTeamid(3, 1)
	teamid, err = GetUserTeamid(3)
	if teamid != 1 {
		t.Fatalf("Expected team 1. Got: '%v' from user 3", teamid)
	}
	t.Logf("Pass: Team successfully updated. Got: '%v' from user 3", teamid)

	// ------- Testing DeleteUser Function -------
	DeleteUser(2)

	// ------- Testing GetAllUsers Function -------
	ids, err := GetAllUsers()
	if err != nil {
		t.Fatalf("Failed to get all users: '%v'", err)
	}
	t.Logf("Pass: Successfully got all user IDS: '%v'", ids)

	// ------- Testing GetUser Function -------
	user1 := GetUser(1)
	t.Logf("Pass: Got user1: '%v'", user1)
}
