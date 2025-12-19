package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "General.html", nil)
	})

	r.GET("/myproject", func(c *gin.Context) {
		c.HTML(200, "MyProject.html", nil)
	})

	r.Run(":8080")
}
