package taskdb

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const DIRECTORY = "../../data/task.db"

func InitDatabase() {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	create_task_db := `
	CREATE TABLE IF NOT EXISTS tasks (
		taskid INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT DEFAULT "",
		duedate TEXT DEFAULT "",
		timeframe TEXT DEFAULT "",
		schedule TEXT DEFAULT "",
		points INTEGER DEFAULT 0,
		triggers INTEGER DEFAULT 0
	);
	`

	_, err = db.Exec(create_task_db)
	if err != nil {
		log.Printf("%q: %s\n", err, create_task_db)
		return
	}
}

func CreateTask(title string) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
		db.Close()
		return
	}
	defer db.Close()

	create_task_sql := `
	INSERT INTO tasks (title)
	VALUES (?)
	`

	_, err = db.Exec(create_task_sql, title)
	if err != nil {
		log.Printf("%q: %s\n", err, create_task_sql)
		return
	}
}

func DeleteTask(taskid int) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
		db.Close()
		return
	}
	defer db.Close()

	delete_task_sql := `
	DELETE FROM tasks WHERE taskid = ?
	`

	_, err = db.Exec(delete_task_sql, taskid)
	if err != nil {
		log.Printf("%q: %s\n", err, delete_task_sql)
		return
	}
}

func GetTaskTitle(taskid int) (string, error) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
		db.Close()
		return "", err
	}
	defer db.Close()

	retrieve_task_sql := `
	SELECT title FROM tasks WHERE taskid = ?
	`

	var title string
	err = db.QueryRow(retrieve_task_sql, taskid).Scan(&title)
	if err != nil {
		log.Printf("%q: %s\n", err, retrieve_task_sql)
		return "", err
	}
	return title, nil
}

func AssignTaskTitle(taskid int, title string) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
		db.Close()
		return
	}
	defer db.Close()

	update_title_sql := `
	UPDATE tasks SET title = ?
	WHERE taskid = ?
	`

	_, err = db.Exec(update_title_sql, title, taskid)
	if err != nil {
		log.Printf("%q: %s\n", err, update_title_sql)
		return
	}
}

func GetTaskDescription(taskid int) (string, error) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
		db.Close()
		return "", err
	}
	defer db.Close()

	retrieve_description_sql := `
	SELECT description FROM tasks WHERE taskid = ?
	`

	var description string
	err = db.QueryRow(retrieve_description_sql, taskid).Scan(&description)
	if err != nil {
		log.Printf("%q: %s\n", err, retrieve_description_sql)
		return "", err
	}
	return description, nil
}

func AssignTaskDescription(taskid int, description string) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
		db.Close()
		return
	}
	defer db.Close()

	update_description_sql := `
	UPDATE tasks SET description = ?
	WHERE taskid = ?
	`

	_, err = db.Exec(update_description_sql, description, taskid)
	if err != nil {
		log.Printf("%q: %s\n", err, update_description_sql)
		return
	}
}

func GetTaskDuedate(taskid int) (string, error) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
		db.Close()
		return "", err
	}
	defer db.Close()

	retrieve_duedate_sql := `
	SELECT duedate FROM tasks WHERE taskid = ?
	`

	var duedate string
	err = db.QueryRow(retrieve_duedate_sql, taskid).Scan(&duedate)
	if err != nil {
		log.Printf("%q: %s\n", err, retrieve_duedate_sql)
		return "", err
	}
	return duedate, nil
}

func AssignTaskDuedate(taskid int, duedate string) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
		db.Close()
		return
	}
	defer db.Close()

	update_duedate_sql := `
	UPDATE tasks SET duedate = ?
	WHERE taskid = ?
	`

	_, err = db.Exec(update_duedate_sql, duedate, taskid)
	if err != nil {
		log.Printf("%q: %s\n", err, update_duedate_sql)
		return
	}
}

func GetTaskTimeframe(taskid int) (string, error) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
		db.Close()
		return "", err
	}
	defer db.Close()

	retrieve_timeframe_sql := `
	SELECT timeframe FROM tasks WHERE taskid = ?
	`

	var timeframe string
	err = db.QueryRow(retrieve_timeframe_sql, taskid).Scan(&timeframe)
	if err != nil {
		log.Printf("%q: %s\n", err, retrieve_timeframe_sql)
		return "", err
	}
	return timeframe, nil
}

func AssignTaskTimeframe(taskid int, timeframe string) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
		db.Close()
		return
	}
	defer db.Close()

	update_timeframe_sql := `
	UPDATE tasks SET timeframe = ?
	WHERE taskid = ?
	`

	_, err = db.Exec(update_timeframe_sql, timeframe, taskid)
	if err != nil {
		log.Printf("%q: %s\n", err, update_timeframe_sql)
		return
	}
}
