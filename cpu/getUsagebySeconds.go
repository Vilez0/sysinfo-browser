package cpu

import (
	"fmt"
	"htop/util"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// "github.com/spf13/viper"
)

type CpuUsage struct {
	Time  int64 `gorm:"primaryKey;autoIncrement:false"`
	Usage string
}

type Usage struct {
	Usage string
}

func GetUsagebySeconds(seconds int) []float64 {
	var cpuUsages []float64
	// database := viper.Get("database")
	db, err := gorm.Open(sqlite.Open("db/usage.db"), &gorm.Config{})
	if err != nil {
		util.ErrorLogger.Printf("Cannot open database: %v\n", err)
	}
	var usages []Usage
	// data := db.First(&usage, "time >= UNIXEPOCH('now','-"+strconv.Itoa(seconds)+" seconds')")
	data := db.Table(`cpu_usages`).Where("time >= ?", time.Now().Unix()-int64(seconds)).Select(`usage`).Find(&usages)
	if data.Error != nil {
		util.ErrorLogger.Printf(`Error when searching database: %v\n`, data.Error)
	}
	println(data.RowsAffected, data.Error, data)
	fmt.Printf("%v\n", usages[0])
	fmt.Print(len(usages))
	// var strUsages []string

	for e := range usages {
		usage := fmt.Sprint(usages[e])
		usage = strings.ReplaceAll(usage, "[", "")
		usage = strings.ReplaceAll(usage, "]", "")
		usage = strings.ReplaceAll(usage, "{", "")
		usage = strings.ReplaceAll(usage, "}", "")
		usages := strings.Split(usage, " ")
		println(usage)
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
