package server

import (
	"github.com/gin-gonic/gin"
)

func ServeIndex(c *gin.Context) {
	file := c.Param("file")
	if file == "" {
		c.File("index.html")
	} else if file == "index.css" || file == "index.mjs" {
		c.File(file)
	}
}
