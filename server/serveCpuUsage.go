package server

import (
	"htop/info/cpu"
	"htop/util"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ServeCpuUsage(c *gin.Context) {
	seconds := c.Param("seconds")
	average := strings.ReplaceAll(c.Param("average"), "/", "")
	//*Check the url, then serve the content as the url
	if seconds == "" {
		usage, _ := cpu.Usage()

		c.String(200, marshaler(usage))
		return
	} else if seconds == "average" && average == "" {
		_, average := cpu.Usage()
		c.String(200, marshaler(average))
		return

	} else if seconds == "average" && average != "" {
		seconds, err := strconv.Atoi(average)
		if err != nil {
			util.ErrorLogger.Printf("error when converting string to integar: %v", err)
			return
		}
		values := cpu.UsagebySeconds(seconds)
		average, _ := cpu.CalculateConfidenceInterval(values)
		c.String(200, marshaler(int(average)))
		return
	} else if seconds == "cinterval" && average != "" {
		seconds, err := strconv.Atoi(average)
		if err != nil {
			util.ErrorLogger.Printf("error when converting string to integar: %v", err)
			return
		}
		values := cpu.UsagebySeconds(seconds)

		_, cinterval := cpu.CalculateConfidenceInterval(values)
		var intcInterval []int
		for _,e := range cinterval {
			intcInterval = append(intcInterval, int(e))
		}
		c.String(200, marshaler(intcInterval))
		return
	} else if seconds != "" && average == "" {
		seconds, err := strconv.Atoi(seconds)
		if err != nil {
			util.ErrorLogger.Printf("error when converting string to integar: %v", err)
			return
		}
		//* Return all values stored in column usage in last $seconds seconds
		c.String(200, marshaler(cpu.UsagebySeconds(seconds)))
	}
}
