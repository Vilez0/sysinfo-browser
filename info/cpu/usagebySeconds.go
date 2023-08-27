package cpu

import (
	"fmt"
	"htop/util"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type CpuUsage struct {
	Time  int64 `gorm:"primaryKey;autoIncrement:false"`
	Usage string
}

type Usages struct {
	Usage string
}

func UsagebySeconds(seconds int) []float64 {
	var cpuUsages []float64
	var usages []Usages
	var data *gorm.DB
	db, err := gorm.Open(sqlite.Open("db/usage.db"), &gorm.Config{})
	if err != nil {
		util.ErrorLogger.Printf("Cannot open database: %v\n", err)
	}
	if seconds == 1 || seconds == 0 {
		data = db.Table(`cpu_usages`).Where("time = ?", time.Now().Unix()-int64(seconds)).Select(`usage`).Find(&usages)
	} else {
		data = db.Table(`cpu_usages`).Where("time >= ?", time.Now().Unix()-int64(seconds)).Select(`usage`).Find(&usages)

	}
	if data.Error != nil {
		util.ErrorLogger.Printf(`Error when searching database: %v\n`, data.Error)
	}

	for e := range usages {
		usage := fmt.Sprint(usages[e])
		replacer := strings.NewReplacer("[", "", "]", "", "{", "", "}", "")
		usage = replacer.Replace(usage)
		usages := strings.Split(usage, " ")

		for element := range usages {
			toFloat64, err := strconv.ParseFloat(usages[element], 32)
			if err != nil {
				util.ErrorLogger.Printf("Cannot convert string to float64: %v\n", err)
			}
			cpuUsages = append(cpuUsages, toFloat64)
		}
	}
	return cpuUsages
}
