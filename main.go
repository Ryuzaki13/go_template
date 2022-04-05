package main

import (
	"github.com/gin-gonic/gin"
)

type Task struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Text   string `json:"text"`
	Status string `json:"status"` // active | complete | delete
}

var tasks []Task

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("template/*.html")

	router.Static("assets", "assets")

	router.GET("/", index)
	router.PUT("/task", addTask)
	router.GET("/task", getTasks)
	router.POST("/task", updateStatus)

	_ = router.Run("127.0.0.1:8080")
}

func index(c *gin.Context) {
	c.HTML(200, "index", nil)
}

func getTasks(c *gin.Context) {
	c.JSON(200, tasks)
}

func updateStatus(c *gin.Context) {
	var task Task

	e := c.BindJSON(&task)
	if e != nil {
		c.JSON(400, nil)
		return
	}

	for i := range tasks {
		if task.ID == tasks[i].ID {
			tasks[i].Status = task.Status
			break
		}
	}

	c.JSON(200, tasks)
}

func addTask(c *gin.Context) {
	var task Task

	e := c.BindJSON(&task)
	if e != nil {
		c.JSON(400, nil)
		return
	}

	task.ID = len(tasks) + 1
	task.Status = "active"

	tasks = append(tasks, task)

	c.JSON(200, tasks)
}

// https://github.com/Ryuzaki13/go_template
