package main

import (
	"awesomeProject2/db"
	"awesomeProject2/setting"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitLogger()

	opt := setting.Load("setting.json")
	e := db.Connect(opt)
	if e != nil {
		fmt.Println(e)
		return
	}

	router := gin.Default()
	router.LoadHTMLGlob("template/*")
	router.Static("assets", "assets")
	router.GET("/", index)
	router.POST("/user", index2)
	_ = router.Run(opt.Address + ":" + opt.Port)
}

func index(c *gin.Context) {
	user := db.User{}

	users := user.SelectAll()

	c.HTML(200, "index", gin.H{
		"Users":   users,
		"Title":   "Сайтик",
		"IsAdmin": true,
	})
}

func index2(c *gin.Context) {

	type requestData struct {
		Date string `json:"Date"`
	}

	var data requestData

	e := c.BindJSON(&data)
	if e != nil {
		fmt.Println(e)
		c.JSON(400, nil)
		return
	}

	fmt.Println(data)

	c.JSON(200, gin.H{
		"Users":   "HELLO",
		"IsAdmin": true,
		"Date":    "2022-02-18",
	})
}
