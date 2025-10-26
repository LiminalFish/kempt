package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// --- Structs to match your API's JSON responses ---
type Task struct {
	TaskID      int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     int64  `json:"duedate"`
	Timeframe   string `json:"timeframe"` // We get this, but don't use it
	Points      int    `json:"points"`
	Triggers    int    `json:"triggers"`
	Hidden      bool   `json:"hidden"`
	Completed   bool
}

type User struct {
	UserID int    `json:"userid"`
	Name   string `json:"name"`
	Type   int    `json:"type"`
	Points int    `json:"points"`
	Rank   int    `json:"rank"`
}

func (u User) Role() string {
	switch u.Type {
	case 1:
		return "Admin"
	case 2:
		return "Moderator"
	case 3:
		return "User"
	default:
		return "Unknown"
	}
}

// --- Configuration ---
const (
	apiBaseUrl      = "http://localhost:8081"
	userIDToDisplay = 1
)

var (
	authToken  = ""
	tokenMutex = &sync.RWMutex{}
)

// --- Template Helper Function ---
func formatAsDate(unixTime int64) string {
	if unixTime == 0 {
		return ""
	}
	t := time.Unix(unixTime, 0)
	return t.Format("2006-01-02") // Format as YYYY-MM-DD for <input type="date">
}

// --- API Client Helper ---
func apiRequest(method, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, apiBaseUrl+endpoint, body)
	if err != nil {
		return nil, err
	}
	tokenMutex.RLock()
	req.Header.Set("auth", authToken)
	tokenMutex.RUnlock()
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 10 * time.Second}
	return client.Do(req)
}

func main() {
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.Static("/static", "./web/static")
	router.LoadHTMLGlob("./web/templates/*")

	// Routes
	router.GET("/", showIndexPage)
	router.GET("/login", showLoginPage)
	router.POST("/login", handleLogin)
	router.GET("/logout", handleLogout)

	// NEW/MODIFIED routes
	router.POST("/add", addTask)
	router.POST("/complete/:id", completeTask)   // Replaces delete
	router.POST("/configure/:id", configureTask) // For the new modal

	log.Println("Frontend server starting on http://localhost:8080")
	router.Run(":8080")
}

// --- Route Handlers ---

func showLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func handleLogin(c *gin.Context) {
	token := c.PostForm("token")
	if token == "" {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"Error": "Token cannot be empty"})
		return
	}
	tokenMutex.Lock()
	authToken = token
	tokenMutex.Unlock()
	log.Println("Auth token stored")
	c.Redirect(http.StatusFound, "/")
}

func handleLogout(c *gin.Context) {
	tokenMutex.Lock()
	authToken = ""
	tokenMutex.Unlock()
	log.Println("User logged out, token cleared")
	c.Redirect(http.StatusFound, "/login")
}

func showIndexPage(c *gin.Context) {
	if authToken == "" {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	var user User
	resp, err := apiRequest("GET", "/users/"+strconv.Itoa(userIDToDisplay), nil)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"Error": "Failed to fetch user: " + err.Error()})
		return
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"Error": "Failed to parse user data: " + err.Error()})
		return
	}

	type TaskIDsResponse struct {
		TaskIDs []int `json:"taskIDs"`
	}
	var taskIDsResp TaskIDsResponse
	resp, err = apiRequest("GET", "/tasks", nil)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{"Error": "Failed to fetch task IDs: " + err.Error(), "User": user})
		return
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&taskIDsResp); err != nil {
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{"Error": "Failed to parse task IDs: " + err.Error(), "User": user})
		return
	}

	var todos []Task
	for _, id := range taskIDsResp.TaskIDs {
		var task Task
		taskResp, err := apiRequest("GET", "/tasks/"+strconv.Itoa(id), nil)
		if err != nil {
			log.Printf("Failed to fetch task %d: %v", id, err)
			continue
		}
		defer taskResp.Body.Close()
		if err := json.NewDecoder(taskResp.Body).Decode(&task); err != nil {
			log.Printf("Failed to parse task %d: %v", id, err)
			continue
		}
		if !task.Hidden {
			todos = append(todos, task)
		}
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"User":  user,
		"Todos": todos,
		"Dones": []Task{},
	})
}

// addTask no longer needs to handle duedate/points, as this is done in "Configure"
func addTask(c *gin.Context) {
	title := c.PostForm("title")

	if title == "" {
		c.Redirect(http.StatusFound, "/")
		return
	}

	// This matches your API's POST /tasks/create endpoint (title only)
	type CreateTaskRequest struct {
		Title string `json:"title"`
	}

	if err != nil {
		log.Println(err.Error())
		return
	}

	reqBody := CreateTaskRequest{Title: title}
	jsonBody, _ := json.Marshal(reqBody)

	resp, err := apiRequest("POST", "/tasks/create", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Printf("Failed to create task: %v", err)
	} else {
		resp.Body.Close()
		log.Printf("API: Created task '%s'", title)
	}
	c.Redirect(http.StatusFound, "/")
}

// completeTask awards points and then deletes the task
func completeTask(c *gin.Context) {
	idStr := c.Param("id")
	// taskID, _ := strconv.Atoi(idStr)

	// --- 1. Get Task Points ---
	var task Task
	resp, err := apiRequest("GET", "/tasks/"+idStr, nil)
	if err != nil {
		log.Printf("CompleteTask: Failed to get task details: %v", err)
		c.Redirect(http.StatusFound, "/")
		return
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&task)
	log.Printf("Task %d is worth %d points", task.TaskID, task.Points)

	// --- 2. Get User's Current Points ---
	var user User
	resp, err = apiRequest("GET", "/users/"+strconv.Itoa(userIDToDisplay), nil)
	if err != nil {
		log.Printf("CompleteTask: Failed to get user details: %v", err)
		c.Redirect(http.StatusFound, "/")
		return
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&user)
	log.Printf("User %d has %d points", user.UserID, user.Points)

	// --- 3. Update User's Points ---
	// Your API (post_user.go) expects: {"newpoint": ...}
	type PointChangeRequest struct {
		NewPoint int `json:"newpoint"`
	}
	newTotalPoints := user.Points + task.Points
	reqBody := PointChangeRequest{NewPoint: newTotalPoints}
	jsonBody, _ := json.Marshal(reqBody)

	resp, err = apiRequest("POST", "/users/"+strconv.Itoa(userIDToDisplay)+"/points", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Printf("CompleteTask: Failed to update user points: %v", err)
		c.Redirect(http.StatusFound, "/")
		return
	}
	resp.Body.Close()
	log.Printf("API: Updated user %d points to %d", user.UserID, newTotalPoints)

	// --- 4. Delete the Task ---
	resp, err = apiRequest("POST", "/tasks/"+idStr+"/delete", nil)
	if err != nil {
		log.Printf("CompleteTask: Failed to delete task: %v", err)
	} else {
		resp.Body.Close()
		log.Printf("API: Deleted task %s", idStr)
	}

	c.Redirect(http.StatusFound, "/")
}

// configureTask handles the modal form submission
func configureTask(c *gin.Context) {
	idStr := c.Param("id")

	// Get all form values
	title := c.PostForm("title")
	description := c.PostForm("description")
	duedateStr := c.PostForm("duedate")
	pointsStr := c.PostForm("points")

	// We have to make multiple API calls, one for each field.
	// We'll run them in parallel for speed.
	var wg sync.WaitGroup

	// 1. Update Title
	wg.Add(1)
	go func() {
		defer wg.Done()
		type TitleChange struct {
			Title string `json:"title"`
		}
		jsonBody, _ := json.Marshal(TitleChange{Title: title})
		resp, err := apiRequest("POST", "/tasks/"+idStr+"/title", bytes.NewBuffer(jsonBody))
		if err == nil {
			resp.Body.Close()
			log.Printf("API: Updated task %s title", idStr)
		}
	}()

	// 2. Update Description
	wg.Add(1)
	go func() {
		defer wg.Done()
		type DescChange struct {
			Description string `json:"description"`
		}
		jsonBody, _ := json.Marshal(DescChange{Description: description})
		resp, err := apiRequest("POST", "/tasks/"+idStr+"/description", bytes.NewBuffer(jsonBody))
		if err == nil {
			resp.Body.Close()
			log.Printf("API: Updated task %s description", idStr)
		}
	}()

	// 3. Update Points
	wg.Add(1)
	go func() {
		defer wg.Done()
		points, _ := strconv.Atoi(pointsStr)
		type PointsChange struct {
			Points int `json:"points"`
		}
		jsonBody, _ := json.Marshal(PointsChange{Points: points})
		resp, err := apiRequest("POST", "/tasks/"+idStr+"/points", bytes.NewBuffer(jsonBody))
		if err == nil {
			resp.Body.Close()
			log.Printf("API: Updated task %s points", idStr)
		}
	}()

	// 4. Update DueDate
	wg.Add(1)
	go func() {
		defer wg.Done()
		var duedateInt int
		if duedateStr != "" {
			t, err := time.Parse("2006-01-02", duedateStr)
			if err == nil {
				duedateInt = int(t.Unix()) // Your API expects an int, not int64
			}
		}
		type DueDateChange struct {
			DueDate int `json:"duedate"`
		}
		jsonBody, _ := json.Marshal(DueDateChange{DueDate: duedateInt})
		resp, err := apiRequest("POST", "/tasks/"+idStr+"/duedate", bytes.NewBuffer(jsonBody))
		if err == nil {
			resp.Body.Close()
			log.Printf("API: Updated task %s duedate", idStr)
		}
	}()

	wg.Wait() // Wait for all API calls to finish
	c.Redirect(http.StatusFound, "/")
}
