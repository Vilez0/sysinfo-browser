package mem

import (
	"strconv"
	"strings"

	"github.com/Edip1/sysinfo-browser/util"
)

func Available() int {
	line := util.ReadProcFile("/proc/meminfo", "MemAvailable:")
	memAvailableKB, _ := strconv.Atoi(strings.Fields(line[0])[1])
	memavailable := memAvailableKB / 1024
	return memavailable
}
