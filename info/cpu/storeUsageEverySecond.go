package cpu

import (
	"htop/util"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func StoreCpuUsageEverySecond() {
	//* Keep executing this function every 1 second
	for range time.Tick(time.Second * 1) {
		db, err := gorm.Open(sqlite.Open("db/usage.db"), &gorm.Config{})
		if err != nil {
			util.ErrorLogger.Printf("Cannot open database: %v", err)
			return
		}
		usage, _ := Usage()
		data := CpuUsage{Time: time.Now().Unix(), Usage: usage}
		err = db.AutoMigrate(&CpuUsage{})
		if err != nil {
			util.ErrorLogger.Printf("Cannot migrate database: %v", err)
			// return
		}
		result := db.Create(&data)
		if result.Error != nil {
			util.ErrorLogger.Println(`Error When creating data: `, result.Error)
		}

		db.Delete(&CpuUsage{}, "time <= ?", time.Now().Unix()-86400)

	}
}
