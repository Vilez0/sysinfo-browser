package server

// import (
// 	"strings"

// 	"github.com/gin-gonic/gin"
// )

// func ServeSystem(ctx *gin.Context) {
// 	name := ctx.Param("name")
// 	info := ctx.Param("info")
// 	info = strings.ReplaceAll(info, "/", "")
// 	info2 := ctx.Param("info2")
// 	info2 = strings.ReplaceAll(info2, "/", "")

// 	if name == "os" {
// 		serveOSInfo(info, ctx)
// 	} else if name == "cpu" {
// 		serverCpuInfo(info, ctx)
// 	} else if name == "gpu" {
// 		serverGpuInfo(info, ctx)
// 	} else if name == "mem" {
// 		serveMemInfo(info, ctx)
// 	} else if name == "disks" {
// 		serveDisks(info, info2, ctx)
// 	}

// }
