package cpu

import (
	"htop/util"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// "github.com/spf13/viper"
)

type CpuUsage struct {
	time  int64  `gorm:"column:TIME;type:int;primaryKey;unique;not null"`
	usage []byte `gorm:"column:USAGE;type:blob;not null;"`
}
type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (CpuUsage) TableName() string {
	return `CPU_USAGE`
}

func GetUsagebySeconds(seconds int) []float64 {
	var cpuUsages []float64
	// database := viper.Get("database")
	db, err := gorm.Open(sqlite.Open("db/usage.db"), &gorm.Config{})
	if err != nil {
		util.ErrorLogger.Printf("Cannot open database: %v", err)
	}

	data := db.First(&CpuUsage{}, "time >= UNIXEPOCH('now','-"+strconv.Itoa(seconds)+" seconds')")
	print(data)
	// db, err := sql.Open("sqlite3", database)
	// //* Get the usage data from the database that stored in last x seconds
	// data, err := db.Query("SELECT usage FROM CPU_USAGE WHERE time >= UNIXEPOCH('now','-" + strconv.Itoa(seconds) + " seconds');")
	// if err != nil {
	// 	ErrorLogger.Printf("Cannot extract usage data from database: %v", err)
	// }

	// //* Scan every element and append it to cpuUsages slice
	// var cpuUsages []float64
	// defer data.Close()
	// for data.Next() {
	// 	var usage string
	// 	err = data.Scan(&usage)
	// 	if err != nil {
	// 		ErrorLogger.Printf("Cannot scan database: %v\n", err)
	// 	}

	// 	//* The data is stored like [value1,value2,value3] in string type, here we are parsing the data to get the values
	// 	usage = strings.Replace(usage, "[", "", -1)
	// 	usage = strings.Replace(usage, "]", "", -1)
	// 	usages := strings.Split(usage, ",")

	// 	//* append the values to slice
	// 	for element := range usages {
	// 		toFloat64, err := strconv.ParseFloat(usages[element], 32)
	// 		if err != nil {
	// 			ErrorLogger.Printf("Cannot convert string to float64: %v\n", err)
	// 		}
	// 		cpuUsages = append(cpuUsages, toFloat64)
	// 	}
	// }
	return cpuUsages
}
