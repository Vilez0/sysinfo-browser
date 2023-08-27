package server

import (
	"htop/info/mem"
	"htop/util"
	"strings"

	"github.com/gin-gonic/gin"
)

func serveMemInfo(ctx *gin.Context) {
	info := ctx.Param("info")
	info = strings.ReplaceAll(info, "/", "")

	switch info {
	case "total":
		ctx.Data(util.ReturnResponse(mem.Total(), 200, "ok", "Memory Size(MiB)"))
	case "usage":
		ctx.Data(util.ReturnResponse(mem.UsageMB(), 200, "ok", "Memory Usage (MiB)"))
	case "usagepercent":
		ctx.Data(util.ReturnResponse(mem.UsagePercent(), 200, "ok", "Memory Usage (Percent)"))
	case "available":
		ctx.Data(util.ReturnResponse(mem.Available(), 200, "ok", "Available Memory(MiB)"))
	}
}
