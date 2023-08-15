package server

import (
	"github.com/gin-gonic/gin"
)

func Run(addr string) {
	router := gin.Default()
	router.GET("/", ServeIndex)
	router.GET("/:file", ServeIndex)
	router.GET("/realtime/cpus/", ServeCpuUsage)
	router.GET("/realtime/cpus/:seconds/*average", ServeCpuUsage)

	router.Run(addr)
}
