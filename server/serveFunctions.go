package server

import (
	"htop/info/cpu"
	"htop/info/disk"
	"htop/info/gpu"
	"htop/info/mem"
	osutils "htop/info/os"
	"htop/util"
	"strings"

	"github.com/gin-gonic/gin"
)

func serveOSInfo(ctx *gin.Context) {
	info := ctx.Param("info")
	info = strings.ReplaceAll(info, "/", "")

	if info == "hostname" {
		ctx.Data(util.ReturnResponse(osutils.Hostname(),200, "ok", "Hostname") )
	} else if info == "name" {
		ctx.Data(util.ReturnResponse(osutils.OsName(),200, "ok", "Distro Name"))
	} else if info == "kernel" {
		ctx.Data(util.ReturnResponse(osutils.KernelName(),200, "ok", "Kernel"))
	} else if info == "desktop" {
		ctx.Data(util.ReturnResponse(osutils.Desktop(),200, "ok", "Desktop Environment"))
	}
}

func serverCpuInfo(ctx *gin.Context) {
	info := ctx.Param("info")
	info = strings.ReplaceAll(info, "/", "")
	if info == "name" {
		ctx.Data(util.ReturnResponse(cpu.Name(),200, "ok", "CPU Name"))
	}
}

func serverGpuInfo(ctx *gin.Context) {
	info := ctx.Param("info")
	info = strings.ReplaceAll(info, "/", "")

	if info == "name" {
		ctx.Data(util.ReturnResponse(gpu.Name(),200, "ok", "GPU Name"))
	}
}

func serveMemInfo(ctx *gin.Context) {
	info := ctx.Param("info")
	info = strings.ReplaceAll(info, "/", "")

	switch info {
	case "total":
		ctx.Data(util.ReturnResponse(mem.Total(),200, "ok", "Memory Size(MiB)"))
	case "usage":
		ctx.Data(util.ReturnResponse(mem.UsageMB(),200, "ok", "Memory Usage (MiB)"))
	case "usagepercent":
		ctx.Data(util.ReturnResponse(mem.UsagePercent(),200, "ok", "Memory Usage (Percent)"))
	case "available":
		ctx.Data(util.ReturnResponse(mem.Available(),200, "ok", "Available Memory(MiB)"))
	}
}

func serveDisks(ctx *gin.Context) {
	part := ctx.Param("part")
	info := ctx.Param("info")
	disks := disk.Disks()
	if part == "" {
		ctx.Data(util.ReturnResponse(disks,200, "ok", "Disks"))
	}
	for _, element := range disks {
		if element == part {
			switch info {
			case "size":
				ctx.Data(util.ReturnResponse(disk.Size(part),200, "ok", "Disk Size(GiB)"))
			}
		}
	}
}
