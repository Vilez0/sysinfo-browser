package server

import (
	osutils "htop/info/os"
	"htop/util"
	"strings"

	"github.com/gin-gonic/gin"
)

func serveOSInfo(ctx *gin.Context) {
	info := ctx.Param("info")
	info = strings.ReplaceAll(info, "/", "")

	if info == "hostname" {
		ctx.Data(util.ReturnResponse(osutils.Hostname(), 200, "ok", "Hostname"))
	} else if info == "name" {
		ctx.Data(util.ReturnResponse(osutils.OsName(), 200, "ok", "Distro Name"))
	} else if info == "kernel" {
		ctx.Data(util.ReturnResponse(osutils.KernelName(), 200, "ok", "Kernel"))
	} else if info == "desktop" {
		ctx.Data(util.ReturnResponse(osutils.Desktop(), 200, "ok", "Desktop Environment"))
	}
}
