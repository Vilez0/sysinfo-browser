package server

import (
	// "encoding/json"
	// "htop/util"
	// "reflect"

	"github.com/gin-gonic/gin"
)

func Run(addr string) {
	router := gin.Default()
	router.GET("/", ServeIndex)
	router.GET("/:file", ServeIndex)
	router.GET("/info/os/:info", serveOSInfo)
	router.GET("/info/cpu/:info", serverCpuInfo)
	router.GET("/info/cpu/:info/:seconds/*average", serverCpuInfo)
	router.GET("/info/gpu/:info", serverGpuInfo)
	router.GET("/info/mem/:info", serveMemInfo)
	router.GET("/info/disks/", serveDisks)
	router.GET("/info/disks/:part/:info/", serveDisks)

	router.Run(addr)
}
