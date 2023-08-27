package main

import (
	"github.com/Edip1/sysinfo-browser/info/cpu"
	"github.com/Edip1/sysinfo-browser/server"
)

func main() {
	go cpu.StoreCpuUsageEverySecond()
	server.Run(":7052")
}
