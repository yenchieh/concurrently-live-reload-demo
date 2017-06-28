package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.Static("/assets", "./dist/assets")
	r.Static("/js", "./dist")
	r.LoadHTMLFiles("./dist/index.html")

	r.GET("/", index)

	r.Run(":8081")
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Message": "This is awesome DEMO",
	})
}