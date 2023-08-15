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
		// ds,err := sql.Open("sqlite3","db/usage.db")
		db, err := gorm.Open(sqlite.Open("db/usage.db"), &gorm.Config{})
		if err != nil {
			util.ErrorLogger.Printf("Cannot open database: %v", err)
			return
		}

		data := CpuUsage{time: time.Now().Unix(), usage: GetUsage()}
		err = db.AutoMigrate(data)
		if err != nil {
			util.ErrorLogger.Printf("Cannot migrate database: %v", err)
			return
		}
		result := db.Create(&data)
		println(result.Error, result.RowsAffected)
		util.ErrorLogger.Println("denneme")
		// db.Delete(&CpuUsage{}, "time <= UNIXEPOCH('now','-24 hours')")

	}
}
