package server

import (
	"htop/info/gpu"
	"htop/util"
	"strings"

	"github.com/gin-gonic/gin"
)

func serverGpuInfo(ctx *gin.Context) {
	info := ctx.Param("info")
	info = strings.ReplaceAll(info, "/", "")

	if info == "name" {
		ctx.Data(util.ReturnResponse(gpu.Name(), 200, "ok", "GPU Name"))
	}
}
