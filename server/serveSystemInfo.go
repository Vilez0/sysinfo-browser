package server

import (
	// cpuInfo "htop/info/cpu"
	osInfo "htop/info/os"
	"strings"

	"github.com/gin-gonic/gin"
)

func ServeSystem(c *gin.Context) {
	name := c.Param("name")
	info := c.Param("info")
	info = strings.ReplaceAll(info, "/", "")
	if name == "os" {
		if info == "hostname" {
			c.String(200, osInfo.Hostname())
		} else if info == "name" {
			c.String(200, osInfo.OsName())
		} else {
			c.String(200, osInfo.KernelName())
		}
	} 
	// else if name == "cpu" {
	// 	if info == "name" {
	// 		c.String(200, cpuInfo.CpuName())
	// 	}
	// }

}
