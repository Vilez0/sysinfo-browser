package server

import (
	"fmt"
	"htop/info/cpu"
	"htop/util"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func serverCpuInfo(ctx *gin.Context) {
	info := ctx.Param("info")
	info = strings.ReplaceAll(info, "/", "")
	seconds := ctx.Param("seconds")
	average := strings.ReplaceAll(ctx.Param("average"), "/", "")
	switch info {
	case "name":
		ctx.Data(util.ReturnResponse(cpu.Name(), 200, "ok", "CPU Name"))
	case "usage":
		realtimeUsage := cpu.UsagebySeconds(1)
		realtimeAverage, _ := cpu.CalculateConfidenceInterval(realtimeUsage)
		switch seconds {
		case "":
			ctx.Data(util.ReturnResponse(realtimeUsage, 200, "ok", "CPU Usage in last 1 second"))
			return
		case "average":
			if average != "" {
				seconds, err := strconv.Atoi(average)
				if err != nil {
					util.ErrorLogger.Printf("error when converting string to integar: %v", err)
					return
				}
				values := cpu.UsagebySeconds(seconds)
				average, _ := cpu.CalculateConfidenceInterval(values)
				ctx.Data(util.ReturnResponse(int(average), 200, "ok", fmt.Sprintf("Average CPU Usage in last %v seconds", seconds)))
				return
			} else {
				ctx.Data(util.ReturnResponse(realtimeAverage, 200, "ok", "Average CPU Usage in last 1 second"))
				return
			}
		case "cinterval":
			if average != "" {
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
			}
		default:
			if average == "" {
				seconds, err := strconv.Atoi(seconds)
				if err != nil {
					util.ErrorLogger.Printf("error when converting string to integar: %v", err)
					return
				}
				//* Return all values stored in column usage in last $seconds seconds
				ctx.Data(util.ReturnResponse(cpu.UsagebySeconds(seconds), 200, "ok", fmt.Sprintf("CPU Usage in last %v seconds", seconds)))
				return
			}
		}
	}
}
