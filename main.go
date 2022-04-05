package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("template/*")

	router.Static("assets", "assets")

	router.GET("/", index)

	_ = router.Run("127.0.0.1:8080")
}

func index(c *gin.Context) {
	c.HTML(200, "index", nil)
}
