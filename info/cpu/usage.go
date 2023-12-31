package cpu

import (
	"strconv"
	"strings"
	"time"

	"github.com/Edip1/sysinfo-browser/util"
)

func getUsageFromIndex(indexNumber int) int {
	var prevIdleTime, prevTotalTime int
	for i := 0; i <= 1; i++ {
		lines := util.ReadProcFile("/proc/stat", "cpu")
		var line string
		if indexNumber > len(lines)-1 {
			return 0
		}

		if indexNumber != 0 {
			line = (lines[indexNumber])[6:] // get rid of cpu plus 2 spaces
		} else {
			line = (lines[indexNumber])[5:] // get rid of cpu0 plus 2 spaces

		}

		split := strings.Fields(line)
		idleTime, _ := strconv.Atoi(split[3])
		totalTime := 0
		for _, s := range split {
			u, _ := strconv.Atoi(s)
			totalTime += u
		}
		if i > 0 {
			deltaIdleTime := idleTime - prevIdleTime
			deltaTotalTime := totalTime - prevTotalTime
			cpuUsage := (1.0 - float64(deltaIdleTime)/float64(deltaTotalTime)) * 100.0
			return int(cpuUsage)
		}
		prevIdleTime = idleTime
		prevTotalTime = totalTime
		time.Sleep(time.Millisecond * 250)
	}
	return 0
}
func Usage() ([]int, int) {
	var CpusUsage []int
	var average int
	for i := range util.ReadProcFile("/proc/stat", "cpu") {
		if i == 0 {
			average = getUsageFromIndex(i)
		} else {
			CpusUsage = append(CpusUsage, getUsageFromIndex(i))
		}
	}

	return CpusUsage, average
}
