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
