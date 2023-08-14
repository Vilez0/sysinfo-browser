package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shirou/gopsutil/cpu"
)

type httpHandlerFunc func(http.ResponseWriter, *http.Request)

// * Define your database location here:
const database string = "../db/usage.db"

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	logFile, err := os.OpenFile("htop.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	InfoLogger = log.New(mw, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(mw, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
func main() {
	go storeCpuUsageEverySecond()
	http.Handle("/", withExecTime(serveIndex))
	http.Handle("/index.mjs", withExecTime(serveIndexjs))
	http.Handle("/index.css", withExecTime(serveIndexCss))
	http.Handle("/realtime/cpus/", withExecTime(serveCpuUsage))
	InfoLogger.Println("starting server")
	err := http.ListenAndServe(":7052", nil)
	if err != nil {
		println(err)
	}

}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func serveIndexCss(w http.ResponseWriter, _ *http.Request) {
	file, err := os.ReadFile("index.css")
	if err != nil {
		ErrorLogger.Printf("error reading index.css file: %v", err)
		return
	}

	w.Header().Set("Content-Type", "text/css;charset=utf-8")
	_, err = w.Write(file)
	if err != nil {
		return
	}
}

func serveIndexjs(w http.ResponseWriter, _ *http.Request) {
	file, err := os.ReadFile("./index.mjs")
	if err != nil {
		ErrorLogger.Printf("error reading index.mjs file: %v", err)
		return
	}

	w.Header().Set("content-type", "application/javascript;charset=utf-8")
	_, err = w.Write(file)
	if err != nil {
		return
	}
}

func serveCpuUsage(w http.ResponseWriter, r *http.Request) {
	url := r.URL.RequestURI()
	//*Check the url, then serve the content as the url
	if url == "/realtime/cpus/" || url == "/realtime/cpus" {
		//* Just return the cpu usage
		fmt.Fprint(w, string(getCpuUsage()))
	} else if strings.Contains(url, "/realtime/cpus/average") {
		//* Grab the seconds from url. if /realtime/cpus/average/x is the url, x will be the seconds
		seconds, err := strconv.Atoi(strings.Split(url, "/realtime/cpus/average/")[1])
		if err != nil {
			ErrorLogger.Printf("error when converting string to integar: %v", err)
			return
		}
		values := getCpuUsageLastSeconds(seconds)
		//* Return the last x seconds cpu usage average and  confidence Interval
		average, confidenceInterval := calculateCpuUsageConfidenceInterval(values)
		fmt.Fprintf(w, "average: %v\nconfidence Interval: %v", average, confidenceInterval)
	} else {
		seconds, err := strconv.Atoi(strings.Split(url, "/realtime/cpus/")[1])
		if err != nil {
			ErrorLogger.Printf("error when converting string to integar: %v", err)
			return
		}
		//* Return all values stored in column usage in last $seconds seconds
		fmt.Fprintf(w, "%v\n", getCpuUsageLastSeconds(seconds))
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
		const createTable string = `
  CREATE TABLE IF NOT EXISTS CPU_USAGE (
  time INTEGER NOT NULL PRIMARY KEY,
  usage BLOB NOT NULL
  );`
		const DeleteLastDay string = `delete FROM CPU_USAGE WHERE time <= UNIXEPOCH('now','-24 hours');`

		cpuUsage := string(getCpuUsage())
		db, err := sql.Open("sqlite3", database)
		if err != nil {
			ErrorLogger.Printf("database error %v", err)
			return
		}

		//* Delete the data that older than 24 hours
		_, err = db.Exec(DeleteLastDay)
		if err != nil {
			ErrorLogger.Printf("database error %v", err)
			return
		}
		//* Create the table if not exists
		_, err = db.Exec(createTable)
		if err != nil {
			ErrorLogger.Printf("database error %v", err)
			return
		}
		//* Append the unixtimestamp and cpu usage into database
		_, err = db.Exec("INSERT INTO CPU_USAGE VALUES(?,?);", time.Now().Unix(), cpuUsage)
		if err != nil {
			ErrorLogger.Printf("database error %v", err)
			return
		}
	}
}

func getCpuUsageLastSeconds(seconds int) []float64 {
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		ErrorLogger.Printf("Cannot open database: %v", err)
	}
	//* Get the usage data from the database that stored in last x seconds
	data, err := db.Query("SELECT usage FROM CPU_USAGE WHERE time >= UNIXEPOCH('now','-" + strconv.Itoa(seconds) + " seconds');")
	if err != nil {
		ErrorLogger.Printf("Cannot extract usage data from database: %v", err)
	}

	//* Scan every element and append it to cpuUsages slice
	var cpuUsages []float64
	defer data.Close()
	for data.Next() {
		var usage string
		err = data.Scan(&usage)
		if err != nil {
			ErrorLogger.Printf("Cannot scan database: %v\n", err)
		}

		//* The data is stored like [value1,value2,value3] in string type, here we are parsing the data to get the values
		usage = strings.Replace(usage, "[", "", -1)
		usage = strings.Replace(usage, "]", "", -1)
		usages := strings.Split(usage, ",")

		//* append the values to slice
		for element := range usages {
			toFloat64, err := strconv.ParseFloat(usages[element], 32)
			if err != nil {
				ErrorLogger.Printf("Cannot convert string to float64: %v\n", err)
			}
			cpuUsages = append(cpuUsages, toFloat64)
		}
	}
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

// * Functions for get the someFunction execute time
func withExecTime(hf httpHandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer duration(time.Now(), r.Method, r.URL)
		hf(w, r)
	})
}
func duration(start time.Time, method string, url *url.URL) {
	execTime := time.Since(start)
	InfoLogger.Printf("%-5s | %-12s | %-20s ", method, execTime, url)
}
