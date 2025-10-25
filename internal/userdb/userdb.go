package userdb

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// initDB creates the local user database if it DOES NOT already exist.
// The user database uses the key 'userid' as the primary.
func InitDatabase() {
	db, err := sql.Open("sqlite3", "./user.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	create_user_db := `
	CREATE TABLE IF NOT EXISTS users (
		userid INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		type INTEGER NOT NULL,
		points INTEGER DEFAULT 0,
		rank INTEGER DEFAULT -1,
		token TEXT NOT NULL,
		teamid INTEGER DEFAULT 0
	);
	`

	_, err = db.Exec(create_user_db)
	if err != nil {
		log.Fatalf("%q: %s\n", err, create_user_db)
		return
	}
}

// Adds a row to users table with provided information.
// Teams, points, rank can be added later.
// 'userid' is automatically created by database
func CreateUser(name string, userType int, token string) {
	db, err := sql.Open("sqlite3", "./user.db")
	if err != nil {
		log.Fatal(err)
		db.Close()
		return
	}
	defer db.Close()

	insert_user_sql := `
	INSERT INTO users (name, type, token)
	VALUES (?, ?, ?)
	`

	_, err = db.Exec(insert_user_sql, name, userType, token)
	if err != nil {
		log.Fatalf("%q: %s\n", err, insert_user_sql)
		return
	}
}

// Retrieves teamid from specified userid.
func GetTeamid(userid int) (int, error) {
	db, err := sql.Open("sqlite3", "./user.db")
	if err != nil {
		log.Fatal(err)
		db.Close()
		return -1, err
	}
	defer db.Close()

	retrieve_team_sql := `
	SELECT teamid FROM users WHERE userid = ?
	`
	var teamid int
	err = db.QueryRow(retrieve_team_sql, userid).Scan(&teamid)
	if err != nil {
		log.Fatalf("%q: %s\n", err, retrieve_team_sql)
		return -1, err
	}
	return teamid, nil
}

// Assign a teamid to a specified userid.
func AssignTeamid(userid int, teamid int) {
	db, err := sql.Open("sqlite3", "./user.db")
	if err != nil {
		log.Fatal(err)
		db.Close()
		return
	}
	defer db.Close()

	update_team_sql := `
	UPDATE users SET teamid = ?
	WHERE userid = ?
	`

	_, err = db.Exec(update_team_sql, teamid, userid)
	if err != nil {
		log.Fatalf("%q: %s\n", err, update_team_sql)
		return
	}
}
