package userdb

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// initDB creates the local user database if it DOES NOT already exist.
// The user database uses the key 'userid' as the primary.
func InitDatabase() {
	db, err := sql.Open("sqlite3", "../../data/user.db")
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
	db, err := sql.Open("sqlite3", "../../data/user.db")
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

func DeleteUser(userid int) {
	db, err := sql.Open("sqlite3", "../../data/user.db")
	if err != nil {
		log.Fatal(err)
		db.Close()
		return
	}
	defer db.Close()

	delete_user_sql := `
	DELETE FROM users WHERE userid = ?
	`

	_, err = db.Exec(delete_user_sql, userid)
	if err != nil {
		log.Fatalf("%q: %s\n", err, delete_user_sql)
		return
	}
}

// Retrieves name from specified userid
func GetUserName(userid int) (string, error) {
	db, err := sql.Open("sqlite3", "../../data/user.db")
	if err != nil {
		log.Fatal(err)
		db.Close()
		return "", err
	}
	defer db.Close()

	retrieve_name_sql := `
	SELECT name FROM users WHERE userid = ?
	`
	var name string
	err = db.QueryRow(retrieve_name_sql, userid).Scan(&name)
	if err != nil {
		log.Fatalf("%q: %s\n", err, retrieve_name_sql)
		return "", err
	}
	return name, nil
}

// Assign a new name to a specified userid.
func AssignUserName(userid int, name string) {
	db, err := sql.Open("sqlite3", "../../data/user.db")
	if err != nil {
		log.Fatal(err)
		db.Close()
		return
	}
	defer db.Close()

	update_name_sql := `
	UPDATE users SET name = ?
	WHERE userid = ?
	`

	_, err = db.Exec(update_name_sql, name, userid)
	if err != nil {
		log.Fatalf("%q: %s\n", err, update_name_sql)
		return
	}
}

// Retrieves user's admin level from specified userid
func GetUserType(userid int) (int, error) {
	db, err := sql.Open("sqlite3", "../../data/user.db")
	if err != nil {
		log.Fatal(err)
		db.Close()
		return -1, err
	}
	defer db.Close()

	retrieve_type_sql := `
	SELECT type FROM users WHERE userid = ?
	`

	var usertype int
	err = db.QueryRow(retrieve_type_sql, userid).Scan(&usertype)
	if err != nil {
		log.Fatalf("%q: %s\n", err, retrieve_type_sql)
		return -1, err
	}
	return usertype, nil
}

// Assign a new admin level to a specified userid.
func AssignUserType(userid int, usertype int) {
	db, err := sql.Open("sqlite3", "../../data/user.db")
	if err != nil {
		log.Fatal(err)
		db.Close()
		return
	}
	defer db.Close()

	update_type_sql := `
	UPDATE users SET type = ?
	WHERE userid = ?
	`

	_, err = db.Exec(update_type_sql, usertype, userid)
	if err != nil {
		log.Fatalf("%q: %s\n", err, update_type_sql)
		return
	}
}

// Retrieves total points from specified userid
func GetUserPoints(userid int) (int, error) {
	db, err := sql.Open("sqlite3", "../../data/user.db")
	if err != nil {
		log.Fatal(err)
		db.Close()
		return -1, err
	}
	defer db.Close()

	retrieve_points_sql := `
	SELECT points FROM users WHERE userid = ?
	`
	var points int
	err = db.QueryRow(retrieve_points_sql, userid).Scan(&points)
	if err != nil {
		log.Fatalf("%q: %s\n", err, retrieve_points_sql)
		return -1, err
	}
	return points, nil
}

// Assign a new value of points to a specified userid.
func AssignUserPoints(userid int, points int) {
	db, err := sql.Open("sqlite3", "../../data/user.db")
	if err != nil {
		log.Fatal(err)
		db.Close()
		return
	}
	defer db.Close()

	update_points_sql := `
	UPDATE users SET points = ?
	WHERE userid = ?
	`

	_, err = db.Exec(update_points_sql, points, userid)
	if err != nil {
		log.Fatalf("%q: %s\n", err, update_points_sql)
		return
	}
}

// Retrieves user rank from specified userid
func GetUserRank(userid int) (int, error) {
	db, err := sql.Open("sqlite3", "../../data/user.db")
	if err != nil {
		log.Fatal(err)
		db.Close()
		return -1, err
	}
	defer db.Close()

	retrieve_rank_sql := `
	SELECT rank FROM users WHERE userid = ?
	`
	var rank int
	err = db.QueryRow(retrieve_rank_sql, userid).Scan(&rank)
	if err != nil {
		log.Fatalf("%q: %s\n", err, retrieve_rank_sql)
		return -1, err
	}
	return rank, nil
}

// Assign a new rank to a specified userid.
func AssignUserRank(userid int, rank int) {
	db, err := sql.Open("sqlite3", "../../data/user.db")
	if err != nil {
		log.Fatal(err)
		db.Close()
		return
	}
	defer db.Close()

	update_rank_sql := `
	UPDATE users SET rank = ?
	WHERE userid = ?
	`

	_, err = db.Exec(update_rank_sql, rank, userid)
	if err != nil {
		log.Fatalf("%q: %s\n", err, update_rank_sql)
		return
	}
}

// Retrieves user API key/token from specified userid
func GetUserToken(userid int) (string, error) {
	db, err := sql.Open("sqlite3", "../../data/user.db")
	if err != nil {
		log.Fatal(err)
		db.Close()
		return "", err
	}
	defer db.Close()

	retrieve_token_sql := `
	SELECT token FROM users WHERE userid = ?
	`
	var token string
	err = db.QueryRow(retrieve_token_sql, userid).Scan(&token)
	if err != nil {
		log.Fatalf("%q: %s\n", err, retrieve_token_sql)
		return "", err
	}
	return token, nil
}

// Retrieves teamid from specified userid.
func GetUserTeamid(userid int) (int, error) {
	db, err := sql.Open("sqlite3", "../../data/user.db")
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
func AssignUserTeamid(userid int, teamid int) {
	db, err := sql.Open("sqlite3", "../../data/user.db")
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
