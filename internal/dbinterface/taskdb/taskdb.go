package taskdb

import (
	"database/sql"
	"log"

	"github.com/LiminalFish/kempt/pkg/models"

	_ "github.com/mattn/go-sqlite3"
)

const DIRECTORY = "./data/task.db"

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
		duedate INTEGER,
		timeframe TEXT DEFAULT "",
		schedule TEXT DEFAULT "",
		points INTEGER DEFAULT 0,
		triggers INTEGER DEFAULT 0,
		hidden BOOLEAN DEFAULT 0
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

func GetTaskDuedate(taskid int) (int, error) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	defer db.Close()

	retrieve_duedate_sql := `
	SELECT duedate FROM tasks WHERE taskid = ?
	`

	var duedate sql.NullInt64
	err = db.QueryRow(retrieve_duedate_sql, taskid).Scan(&duedate)
	if err != nil {
		log.Printf("%q: %s\n", err, retrieve_duedate_sql)
		return -1, err
	}
	if !duedate.Valid {
		return -1, nil
	}
	return int(duedate.Int64), nil
}

func AssignTaskDuedate(taskid int, duedate int) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
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

func GetTaskPoints(taskid int) (int, error) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	defer db.Close()

	retrieve_points_sql := `
	SELECT points FROM tasks WHERE taskid = ?
	`

	var points int
	err = db.QueryRow(retrieve_points_sql, taskid).Scan(&points)
	if err != nil {
		log.Printf("%q: %s\n", err, retrieve_points_sql)
		return -1, err
	}
	return points, nil
}

func AssignTaskPoints(taskid int, points int) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	update_points_sql := `
	UPDATE tasks SET points = ?
	WHERE taskid = ?
	`

	_, err = db.Exec(update_points_sql, points, taskid)
	if err != nil {
		log.Printf("%q: %s\n", err, update_points_sql)
		return
	}
}

func GetTaskTriggers(taskid int) (int, error) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	defer db.Close()

	retrieve_triggers_sql := `
	SELECT triggers FROM tasks WHERE taskid = ?
	`

	var triggers int
	err = db.QueryRow(retrieve_triggers_sql, taskid).Scan(&triggers)
	if err != nil {
		log.Printf("%q: %s\n", err, retrieve_triggers_sql)
		return -1, err
	}
	return triggers, nil
}

func AssignTaskTriggers(taskid int, triggers int) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	update_triggers_sql := `
	UPDATE tasks SET triggers = ?
	WHERE taskid = ?
	`

	_, err = db.Exec(update_triggers_sql, triggers, taskid)
	if err != nil {
		log.Printf("%q: %s\n", err, update_triggers_sql)
		return
	}
}

func GetTaskHiddenStatus(taskid int) (bool, error) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
		return false, err
	}
	defer db.Close()

	retrieve_hidden_sql := `
	SELECT hidden FROM tasks WHERE taskid = ?
	`

	var hidden bool
	err = db.QueryRow(retrieve_hidden_sql, taskid).Scan(&hidden)
	if err != nil {
		log.Printf("%q: %s\n", err, retrieve_hidden_sql)
		return false, err
	}
	return hidden, nil
}

func AssignTaskHiddenStatus(taskid int, hidden bool) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	update_hidden_sql := `
	UPDATE tasks SET hidden = ?
	WHERE taskid = ?
	`

	_, err = db.Exec(update_hidden_sql, hidden, taskid)
	if err != nil {
		log.Printf("%q: %s\n", err, update_hidden_sql)
		return
	}
}

func GetTask(taskid int) models.Task {
	task := &models.Task{}
	task.TaskID = taskid
	task.Title, _ = GetTaskTitle(taskid)
	task.Description, _ = GetTaskDescription(taskid)
	task.DueDate, _ = GetTaskDuedate(taskid)
	task.Timeframe, _ = GetTaskTimeframe(taskid)
	task.Schedule = "" //todo: add scheduling
	task.Points, _ = GetTaskPoints(taskid)
	//todo: add capability for multiple triggers
	task.Triggers, _ = GetTaskTriggers(taskid)
	task.Hidden, _ = GetTaskHiddenStatus(taskid)

	return *task
}

func GetAllTaskIDS() ([]int, error) {
	db, err := sql.Open("sqlite3", DIRECTORY)
	var tasks []int
	if err != nil {
		log.Println(err)
		return tasks, err
	}
	defer db.Close()

	retrieve_allids_sql := `
	SELECT taskid FROM tasks
	`

	rows, err := db.Query(retrieve_allids_sql)
	if err != nil {
		log.Printf("%q: %s\n", err, retrieve_allids_sql)
		return tasks, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			log.Printf("Error scanning row: %v", err)
			return tasks, err
		}
		tasks = append(tasks, id)
	}
	if err = rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v", err)
		return tasks, err
	}
	return tasks, nil
}
