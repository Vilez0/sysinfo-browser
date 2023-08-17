package server

import (
	"htop/info/cpu"
	"htop/info/gpu"
	"htop/info/mem"
	osutils "htop/info/os"
	"strings"

	"github.com/gin-gonic/gin"
)

func ServeSystem(c *gin.Context) {
	name := c.Param("name")
	info := c.Param("info")
	info = strings.ReplaceAll(info, "/", "")
	if name == "os" {
		if info == "hostname" {
			c.String(200, osutils.Hostname())
		} else if info == "name" {
			c.String(200, osutils.OsName())
		} else {
			c.String(200, osutils.KernelName())
		}
	} else if name == "cpu" {
		if info == "name" {
			c.String(200, cpu.Name())
		}
	} else if name == "gpu" {
		if info == "name" {
			c.String(200, "%v", gpu.Name())
		}
	} else if name == "mem" {
		if info == "total" {
			c.String(200, "%v", mem.Total())
		} else if info == "usage" {
			c.String(200, "%v", mem.UsageMB())
		} else if info == "usagepercent" {
			c.String(200, "%v", mem.UsagePercent())
		} else if info == "available" {
			c.String(200, "%v", mem.Available())
		}
	}

}
