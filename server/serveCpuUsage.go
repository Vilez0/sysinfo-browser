package server

import (
	"htop/info/cpu"
	"htop/util"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ServeCpuUsage(ctx *gin.Context) {
	seconds := ctx.Param("seconds")
	average := strings.ReplaceAll(ctx.Param("average"), "/", "")
	//*Check the url, then serve the content as the url
	if seconds == "" {
		usage, _ := cpu.Usage()
		ctx.Data(util.ReturnResponse(usage, 200, "ok", "CPU Usage in last 1 second"))
		return
	} else if seconds == "average" && average == "" {
		_, average := cpu.Usage()
		ctx.Data(util.ReturnResponse(average, 200, "ok", "Average CPU Usage in last 1 second"))
		return

	} else if seconds == "average" && average != "" {
		seconds, err := strconv.Atoi(average)
		if err != nil {
			util.ErrorLogger.Printf("error when converting string to integar: %v", err)
			return
		}
		values := cpu.UsagebySeconds(seconds)
		average, _ := cpu.CalculateConfidenceInterval(values)
		ctx.Data(util.ReturnResponse(int(average), 200, "ok", "Average CPU Usage in last xxx seconds"))
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
		for _, e := range cinterval {
			intcInterval = append(intcInterval, int(e))
		}
		ctx.Data(util.ReturnResponse(intcInterval, 200, "ok", "confidence interval"))
		return
	} else if seconds != "" && average == "" {
		seconds, err := strconv.Atoi(seconds)
		if err != nil {
			util.ErrorLogger.Printf("error when converting string to integar: %v", err)
			return
		}
		//* Return all values stored in column usage in last $seconds seconds
		ctx.Data(util.ReturnResponse(cpu.UsagebySeconds(seconds), 200, "ok", "CPU Usage in last xxx seconds"))
	}
}
