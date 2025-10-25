package taskdb

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase() {
	db, err := sql.Open("sqlite3", "../../data/task.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	create_task_db := `
	CREATE TABLE IF NOT EXISTS tasks (
		taskid INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT DEFAULT "",
		timeframe TEXT DEFAULT "",
		schedule TEXT DEFAULT "",
		points INTEGER DEFAULT 0,
		triggers INTEGER DEFAULT 0
	);
	`

	_, err = db.Exec(create_task_db)
	if err != nil {
		log.Fatalf("%q: %s\n", err, create_task_db)
		return
	}
}

func CreateTask(title string) {
	db, err := sql.Open("sqlite3", "../../data/task.db")
	if err != nil {
		log.Fatal(err)
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
		log.Fatalf("%q: %s\n", err, create_task_sql)
		return
	}
}

func DeleteTask(taskid int) {
	db, err := sql.Open("sqlite3", "../../data/task.db")
	if err != nil {
		log.Fatal(err)
		db.Close()
		return
	}
	defer db.Close()

	delete_task_sql := `
	DELETE FROM tasks WHERE taskid = ?
	`

	_, err = db.Exec(delete_task_sql, taskid)
	if err != nil {
		log.Fatalf("%q: %s\n", err, delete_task_sql)
		return
	}
}

func GetTaskTitle(taskid int) (string, error) {
	db, err := sql.Open("sqlite3", "../../data/task.db")
	if err != nil {
		log.Fatal(err)
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
		log.Fatalf("%q: %s\n", err, retrieve_task_sql)
		return "", err
	}
	return title, nil
}

func AssignTaskTitle(taskid int, title string) {
	db, err := sql.Open("sqlite3", "../../data/task.db")
	if err != nil {
		log.Fatal(err)
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
		log.Fatalf("%q: %s\n", err, update_title_sql)
		return
	}
}

func GetTaskDescription(taskid int) (string, error) {
	db, err := sql.Open("sqlite3", "../../data/task.db")
	if err != nil {
		log.Fatal(err)
		db.Close()
		return "", err
	}
	defer db.Close()

	retrieve_task_sql := `
	SELECT description FROM tasks WHERE taskid = ?
	`

	var title string
	err = db.QueryRow(retrieve_task_sql, taskid).Scan(&title)
	if err != nil {
		log.Fatalf("%q: %s\n", err, retrieve_task_sql)
		return "", err
	}
	return title, nil
}

func AssignTaskDescription(taskid int, title string) {
	db, err := sql.Open("sqlite3", "../../data/task.db")
	if err != nil {
		log.Fatal(err)
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
		log.Fatalf("%q: %s\n", err, update_title_sql)
		return
	}
}
