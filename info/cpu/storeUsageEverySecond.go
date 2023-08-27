package cpu

import (
	"fmt"
	"htop/util"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func StoreCpuUsageEverySecond() {
	// Keep execute this function every 1 second
	for range time.Tick(time.Second * 1) {
		// Open the database
		db, err := gorm.Open(sqlite.Open("db/usage.db"), &gorm.Config{})
		if err != nil {
			util.ErrorLogger.Printf("Cannot open database: %v", err)
			return
		}
		// Get the CPU usage
		usage, _ := Usage()
		// Create a new CpuUsage struct
		data := CpuUsage{Time: time.Now().Unix(), Usage:fmt.Sprint(usage)}
		err = db.AutoMigrate(&CpuUsage{})
		if err != nil {
			util.ErrorLogger.Printf("Cannot migrate database: %v", err)
			// return
		}
		// Create the data
		result := db.Create(&data)
		if result.Error != nil {
			util.ErrorLogger.Println(`Error When creating data: `, result.Error)
		}
		// Delete the data that is older than 1 day
		db.Delete(&CpuUsage{}, "time <= ?", time.Now().Unix()-86400)

	}
}
