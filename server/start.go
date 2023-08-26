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
	router.GET("/realtime/cpus/", ServeCpuUsage)
	router.GET("/realtime/cpus/:seconds/*average", ServeCpuUsage)
	// router.GET("/system/:name/:info/*info2", ServeSystem)

	router.GET("/system/os/:info", serveOSInfo)
	router.GET("/system/cpu/:info", serverCpuInfo)
	router.GET("/system/gpu/:info", serverGpuInfo)
	router.GET("/system/mem/:info", serveMemInfo)
	router.GET("/system/disks/", serveDisks)
	router.GET("/system/disks/:part/:info/", serveDisks)

	router.Run(addr)
}
