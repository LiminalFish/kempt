package actiondb

import (
	"database/sql"
	"log"
	"time"

	"github.com/LiminalFish/kempt/pkg/models"
	_ "github.com/mattn/go-sqlite3"
)

const DIRECTORY = "../../data/action.db"

func InitDatabase() {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	create_user_db := `
	CREATE TABLE IF NOT EXISTS actions (
		timestamp INTEGER NOT NULL,
		userid INT NOT NULL,
		taskname STRING NOT NULL,
		points INTEGER NOT NULL
	);
	`

	_, err = db.Exec(create_user_db)
	if err != nil {
		log.Printf("%q: %s\n", err, create_user_db)
		return
	}
}

func LogAction(userid int, taskname string, points int) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
		db.Close()
		return
	}
	defer db.Close()

	log_action_sql := `
	INSERT INTO actions (timestamp, userid, taskname, points)
	VALUES (?, ?, ?, ?)
	`

	timestamp := time.Now().UnixMilli()

	_, err = db.Exec(log_action_sql, timestamp, userid, taskname, points)
	if err != nil {
		log.Printf("%q: %s\n", err, log_action_sql)
		return
	}
}

func GetActionPoints(timestamp int) (int, error) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
		db.Close()
		return -1, err
	}
	defer db.Close()

	retrieve_points_sql := `
	SELECT points FROM actions WHERE timestamp = ?
	`
	var points int
	err = db.QueryRow(retrieve_points_sql, timestamp).Scan(&points)
	if err != nil {
		log.Printf("%q: %s\n", err, retrieve_points_sql)
		return -1, err
	}
	return points, nil
}

func GetActionTaskname(timestamp int) (string, error) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
		db.Close()
		return "", err
	}
	defer db.Close()

	retrieve_taskname_sql := `
	SELECT taskname FROM actions WHERE timestamp = ?
	`
	var taskname string
	err = db.QueryRow(retrieve_taskname_sql, timestamp).Scan(&taskname)
	if err != nil {
		log.Printf("%q: %s\n", err, retrieve_taskname_sql)
		return "", err
	}
	return taskname, nil
}

func GetActionUserID(timestamp int) (int, error) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
		db.Close()
		return -1, err
	}
	defer db.Close()

	retrieve_userid_sql := `
	SELECT userid FROM actions WHERE timestamp = ?
	`
	var userid int
	err = db.QueryRow(retrieve_userid_sql, timestamp).Scan(&userid)
	if err != nil {
		log.Printf("%q: %s\n", err, retrieve_userid_sql)
		return -1, err
	}
	return userid, nil
}

func GetAction(timestamp int) (models.Action, error) {
	action := &models.Action{}
	action.Timestamp = timestamp
	action.Points, _ = GetActionPoints(timestamp)
	action.TaskName, _ = GetActionTaskname(timestamp)
	action.UserID, _ = GetActionUserID(timestamp)
	return *action, nil
}

func GetAllTimestamps() ([]int, error) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	var timestamps []int
	if err != nil {
		log.Println(err)
		db.Close()
		return timestamps, err
	}
	defer db.Close()

	retrieve_alltimestamps_sql := `
	SELECT timestamp FROM actions
	`

	rows, err := db.Query(retrieve_alltimestamps_sql)
	if err != nil {
		log.Printf("%q: %s\n", err, retrieve_alltimestamps_sql)
		return timestamps, err
	}
	defer rows.Close()
	for rows.Next() {
		var timestamp int
		if err := rows.Scan(&timestamp); err != nil {
			log.Printf("Error scanning row: %v", err)
			return timestamps, err
		}
		timestamps = append(timestamps, timestamp)
	}
	if err = rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v", err)
		return timestamps, err
	}
	return timestamps, nil
}
