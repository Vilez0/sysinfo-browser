package main

import (
	"encoding/json"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// * Define your database location here:
const database string = "db/usage.db"

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

var (
	ErrorLogger *log.Logger
)

func init() {
	logFile, err := os.OpenFile("htop.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	ErrorLogger = log.New(mw, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
func main() {
	router := gin.Default()
	router.GET("/", serveIndex)
	router.GET("/:file", serveIndex)
	router.GET("/realtime/cpus/", serveCpuUsage)
	router.GET("/realtime/cpus/:seconds/*average", serveCpuUsage)
	go storeCpuUsageEverySecond()

	router.Run(":7052")

}

func serveIndex(c *gin.Context) {
	file := c.Param("file")
	println(file)
	if file == "" {
		c.File("index.html")
	} else if file == "index.css" || file == "index.mjs" {
		c.File(file)
	}

}

func serveCpuUsage(c *gin.Context) {
	seconds := c.Param("seconds")
	average := strings.ReplaceAll(c.Param("average"), "/", "")
	//*Check the url, then serve the content as the url
	if seconds == "" {
		c.String(200, string(getCpuUsage()))
		return
	} else if seconds == "average" && average != "" {
		seconds, err := strconv.Atoi(average)
		if err != nil {
			ErrorLogger.Printf("error when converting string to integar: %v", err)
			return
		}
		values := getCpuUsageLastSeconds(seconds)
		//* Return the last x seconds cpu usage average and  confidence Interval
		average, confidenceInterval := calculateCpuUsageConfidenceInterval(values)
		c.String(200, "average: %v\nconfidence Interval: %v", average, confidenceInterval)
		return
	} else {
		seconds, err := strconv.Atoi(seconds)
		if err != nil {
			ErrorLogger.Printf("error when converting string to integar: %v", err)
			return
		}
		//* Return all values stored in column usage in last $seconds seconds
		c.String(200, "%v\n", getCpuUsageLastSeconds(seconds))
	}
}

func getCpuUsage() []byte {
	percent, err := cpu.Percent(0, true)
	if err != nil {
		ErrorLogger.Printf("error getting cpu usage percent: %v", err)
	}
	//* Encode the cpu usage as json
	e, err := json.Marshal(percent)
	if err != nil {
		ErrorLogger.Printf("error marshaling json: %v", err)
	}
	return e
}

func storeCpuUsageEverySecond() {
	//* Keep executing this function every 1 second
	for range time.Tick(time.Second * 1) {
		// ds,err := sql.Open("sqlite3","db/usage.db")
		db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
		if err != nil {
			ErrorLogger.Printf("Cannot open database: %v", err)
			return
		}

		data := CpuUsage{time: time.Now().Unix(), usage: getCpuUsage()}
		err = db.AutoMigrate(data)
		if err != nil {
			ErrorLogger.Printf("Cannot migrate database: %v", err)
			return
		}
		result := db.Create(&data)
		println(result.Error, result.RowsAffected)
		// db.Delete(&CpuUsage{}, "time <= UNIXEPOCH('now','-24 hours')")

	}
}

func getCpuUsageLastSeconds(seconds int) []float64 {
	var cpuUsages []float64
	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		ErrorLogger.Printf("Cannot open database: %v", err)
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

func calculateCpuUsageConfidenceInterval(samples []float64) (float64, []float64) {
	arrayLong := len(samples)
	var sum, sDeviation, confidenceLevel float64
	confidenceLevel = 0.95
	//* Calculating samples mean
	for i := 0; i < arrayLong; i++ {
		sum += (samples[i])
	}
	mean := sum / float64(arrayLong)

	//* Calculating the standard deviation
	for j := 0; j < len(samples); j++ {
		sDeviation += math.Pow(samples[j]-mean, 2)
	}
	sDeviation = math.Sqrt(sDeviation / 10)

	//* Calculate confidence interval and return
	s := (confidenceLevel * (sDeviation / math.Sqrt(float64(arrayLong))))
	highest := mean + s
	lowest := mean - s

	result := []float64{lowest, highest}
	// fmt.Printf("lowest: %v, highest: %v\n", lowest, highest,)
	return mean, result
}