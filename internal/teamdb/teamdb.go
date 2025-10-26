package teamdb

import (
	"database/sql"
	"log"

	"github.com/LiminalFish/kempt/pkg/models"
	_ "github.com/mattn/go-sqlite3"
)

const DIRECTORY = "../../data/team.db"

func InitDatabase() {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	create_user_db := `
	CREATE TABLE IF NOT EXISTS teams (
		teamid INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name STRING NOT NULL
	);
	`

	_, err = db.Exec(create_user_db)
	if err != nil {
		log.Printf("%q: %s\n", err, create_user_db)
		return
	}
}

func CreateTeam(name string) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
		db.Close()
		return
	}
	defer db.Close()

	insert_team_sql := `
	INSERT INTO teams (name)
	VALUES (?)
	`

	_, err = db.Exec(insert_team_sql, name)
	if err != nil {
		log.Printf("%q: %s\n", err, insert_team_sql)
		return
	}
}

func DeleteTeam(teamid int) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
		db.Close()
		return
	}
	defer db.Close()

	delete_team_sql := `
	DELETE FROM teams WHERE teamid = ?
	`

	_, err = db.Exec(delete_team_sql, teamid)
	if err != nil {
		log.Printf("%q: %s\n", err, delete_team_sql)
		return
	}
}

// Retrieves team name from specified teamid
func GetTeamName(teamid int) (string, error) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
		db.Close()
		return "", err
	}
	defer db.Close()

	retrieve_teamname_sql := `
	SELECT name FROM teams WHERE teamid = ?
	`
	var name string
	err = db.QueryRow(retrieve_teamname_sql, teamid).Scan(&name)
	if err != nil {
		log.Printf("%q: %s\n", err, retrieve_teamname_sql)
		return "", err
	}
	return name, nil
}

// Assign a new team name to a specified teamid.
func AssignTeamName(teamid int, name string) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
		db.Close()
		return
	}
	defer db.Close()

	update_teamname_sql := `
	UPDATE teams SET name = ?
	WHERE teamid = ?
	`

	_, err = db.Exec(update_teamname_sql, name, teamid)
	if err != nil {
		log.Printf("%q: %s\n", err, update_teamname_sql)
		return
	}
}

func GetTeam(teamid int) models.Team {
	team := &models.Team{}
	team.Name, _ = GetTeamName(teamid)
	team.TeamID = teamid

	return *team
}

func GetAllTeamIDS() ([]int, error) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	var teams []int
	if err != nil {
		log.Println(err)
		db.Close()
		return teams, err
	}
	defer db.Close()

	retrieve_allids_sql := `
	SELECT teamid FROM teams
	`

	rows, err := db.Query(retrieve_allids_sql)
	if err != nil {
		log.Printf("%q: %s\n", err, retrieve_allids_sql)
		return teams, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			log.Printf("Error scanning row: %v", err)
			return teams, err
		}
		teams = append(teams, id)
	}
	if err = rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v", err)
		return teams, err
	}
	return teams, nil
}
