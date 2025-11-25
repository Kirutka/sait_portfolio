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

	r.GET("/about_me", func(c *gin.Context) {
		c.HTML(200, "About_me.html", nil)
	})

	r.GET("/contacts", func(c *gin.Context) {
		c.HTML(200, "Contacts.html", nil)
	})

	r.Run(":8080")
}
