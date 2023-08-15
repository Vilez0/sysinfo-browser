package main

import (
	"htop/cpu"
	"htop/server"
)

// * Define your database location here:

func main() {
	go cpu.StoreCpuUsageEverySecond()
	server.Run(":7052")
}
