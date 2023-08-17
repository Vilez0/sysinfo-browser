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
		c.String(200, usage)
		return
	} else if seconds == "average" && average == "" {
		_, average := cpu.Usage()
		result := strconv.Itoa(average)
		c.String(200, result)
		return

	} else if seconds == "average" && average != "" {
		seconds, err := strconv.Atoi(average)
		if err != nil {
			util.ErrorLogger.Printf("error when converting string to integar: %v", err)
			return
		}
		values := cpu.UsagebySeconds(seconds)
		//* Return the last x seconds cpu usage average and  confidence Interval
		average, confidenceInterval := cpu.CalculateConfidenceInterval(values)
		c.String(200, "average: %v\nconfidence Interval: %v", average, confidenceInterval)
		return
	} else {
		seconds, err := strconv.Atoi(seconds)
		if err != nil {
			util.ErrorLogger.Printf("error when converting string to integar: %v", err)
			return
		}
		//* Return all values stored in column usage in last $seconds seconds
		c.String(200, "%v\n", cpu.UsagebySeconds(seconds))
	}
}
