package cpu

import (
	"fmt"
	"htop/util"

	"github.com/shirou/gopsutil/cpu"
)

func GetUsage() string {
	percent, err := cpu.Percent(0, true)
	if err != nil {
		util.ErrorLogger.Printf("error getting cpu usage percent: %v", err)
	}
	var intPercent []int
	for e := range percent {
		intPercent = append(intPercent, int(percent[e]))
	}

	result := fmt.Sprint(intPercent)

	return result
}
