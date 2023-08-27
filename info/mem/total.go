package mem

import (
	"strconv"
	"strings"

	"github.com/Edip1/sysinfo-browser/util"
)

func Total() int {
	line := util.ReadProcFile("/proc/meminfo", "MemTotal:")
	memTotalKB, _ := strconv.Atoi(strings.Fields(line[0])[1])
	memtotal := memTotalKB / 1024
	return memtotal
}
