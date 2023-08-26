package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	// Create a new Gin router with default middleware
	r := gin.Default()

	r.GET("/tasks", getTasks)
	r.POST("/tasks", createTask)
	r.PUT("/tasks/:id", updateTask)
	r.DELETE("/tasks/:id", deleteTask)

	r.Run(":8080") // Listen on port 8080
}

// making task type
type task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"duedate"`
	Completed   bool      `json:"complete"`
}

// slice memory
var tasks = []task{
	{
		ID:          "1",
		Title:       "AWS",
		Description: "test it on AWS",
		DueDate:     time.Date(2023, 8, 27, 0, 0, 0, 0, time.UTC),
		Completed:   true,
	},
	{
		ID:          "2",
		Title:       "Blender",
		Description: "continue bleander retopology lap to bum-bum",
		DueDate:     time.Date(2023, 8, 28, 0, 0, 0, 0, time.UTC),
		Completed:   false,
	},
	{
		ID:          "3",
		Title:       "To-Do-list API",
		Description: "Develop the To-Do list API",
		DueDate:     time.Date(2023, 8, 29, 0, 0, 0, 0, time.UTC),
		Completed:   true,
	},
}

// defining C.R.U.D functions Create. Read. Update. Delete

func getTasks(c *gin.Context) {
	// Retrieve task from your data storage and return them as json
	c.IndentedJSON(http.StatusOK, tasks)
}

func createTask(c *gin.Context) {
	// Parse the request body and create a new task
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)

}

func updateTask(c *gin.Context) {
	// Parse the request body and update the specified task
	taskID := c.Param("id")

	var updateTask task
	if err := c.ShouldBindJSON(&updateTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//find and update the task in the tasks slice "logic"
	for i, t := range tasks {
		if t.ID == taskID {
			tasks[i] = updateTask
			c.JSON(http.StatusOK, tasks[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})

}

func deleteTask(c *gin.Context) {
	// Delete the specified task
	taskID := c.Param("id")

	//find and delete the task in the tasks slice "logic"

	for i, t := range tasks {
		if t.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
}
