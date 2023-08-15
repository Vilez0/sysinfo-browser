package cpu

import (
	"encoding/json"
	"htop/util"

	"github.com/shirou/gopsutil/cpu"
)

func GetUsage() []byte {
	percent, err := cpu.Percent(0, true)
	if err != nil {
		util.ErrorLogger.Printf("error getting cpu usage percent: %v", err)
	}
	//* Encode the cpu usage as json
	e, err := json.Marshal(percent)
	if err != nil {
		util.ErrorLogger.Printf("error marshaling json: %v", err)
	}
	return e
}
