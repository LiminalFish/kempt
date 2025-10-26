package teamdb

import (
	"database/sql"
	"log"
)

const DIRECTORY = "../../data/team.db"

func InitDatabase() {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
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
		log.Printf("%q: %s\n", err, create_user_db)
		return
	}
}
