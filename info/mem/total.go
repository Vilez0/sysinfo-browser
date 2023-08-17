package mem

import (
	"htop/util"
	"strconv"
	"strings"
)

func Total() int {
	line := util.ReadProcFile("/proc/meminfo", "MemTotal:")
	memTotalKB, _ := strconv.Atoi(strings.Fields(line[0])[1])
	memtotal := memTotalKB / 1024
	return memtotal
}
