package models

type LoginHeader struct {
	Pass string `json:"pass"`
}

type Task struct {
	TaskID      int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"duedate"`
	Timeframe   string `json:"timeframe"`
	Schedule    string `json:"schedule"`
	Points      int    `json:"points"`
	Triggers    int    `json:"triggers"`
	Hidden      bool   `json:"hidden"`
}

type User struct {
	UserID int    `json:"userid"`
	Name   string `json:"name"`
	Type   int    `json:"type"`
	Points int    `json:"points"`
	Rank   int    `json:"rank"`
	Token  string `json:"token"`
	TeamID int    `json:"teamid"`
}

type Team struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type Action struct {
	Type      string `json:"type"`
	UserID    string `json:"userid"`
	Timestamp int    `json:"timestamp"`
}
