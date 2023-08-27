package server

import (
	"htop/info/disk"
	"htop/util"

	"github.com/gin-gonic/gin"
)

func serveDisks(ctx *gin.Context) {
	part := ctx.Param("part")
	info := ctx.Param("info")
	disks := disk.Disks()
	if part == "" {
		ctx.Data(util.ReturnResponse(disks, 200, "ok", "Disks"))
	}
	for _, element := range disks {
		if element == part {
			switch info {
			case "size":
				ctx.Data(util.ReturnResponse(disk.Size(part), 200, "ok", "Disk Size(GiB)"))
			}
		}
	}
}
