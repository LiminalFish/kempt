package api

type LoginHeader struct {
	Pass string `json:"pass"`
}

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Timeframe   string `json:"timeframe"`
	Type        string `json:"type"`
	Points      int    `json:"points"`
	Triggers    *Task  `json:"triggers"`
	ID          string `json:"id"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Points int    `json:"points"`
	Rank   int    `json:"rank"`
	Token  string `json:"token"`
	ID     string `json:"id"`
	TeamID string `json:"teamid"`
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

