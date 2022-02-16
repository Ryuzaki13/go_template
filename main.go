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

	user := db.User{}
	user.Login = "ASD222"
	user.Password = "123"
	user.Name = "HELLO"
	user.Insert()

	router := gin.Default()
	router.LoadHTMLGlob("template/*")
	router.Static("assets", "assets")
	router.GET("/", index)
	_ = router.Run(opt.Address + ":" + opt.Port)
}

func index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
