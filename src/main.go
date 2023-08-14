package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

	if url == "/realtime/cpus/" || url == "/realtime/cpus" {
		fmt.Fprint(w, string(getCpuUsage()))
	} else {
		seconds, err := strconv.Atoi(strings.Split(url, "/realtime/cpus/")[1])
		if err != nil {
			ErrorLogger.Printf("error when converting string to integar: %v", err)
			return
		}
		fmt.Fprintf(w, "%v\n", getCpuUsageLastSeconds(seconds))
	}
}

func getCpuUsage() []byte {
	percent, err := cpu.Percent(0, true)
	if err != nil {
		ErrorLogger.Printf("error getting cpu percent: %v", err)
	}

	e, err := json.Marshal(percent)
	if err != nil {
		ErrorLogger.Printf("error marshaling json: %v", err)
	}
	return e
}

func storeCpuUsageEverySecond() {
	for range time.Tick(time.Second * 1) {
		const createTable string = `
  CREATE TABLE IF NOT EXISTS CPU_USAGE (
  time INTEGER NOT NULL PRIMARY KEY,
  usage BLOB NOT NULL
  );`
		cpuUsage := string(getCpuUsage())
		db, err := sql.Open("sqlite3", database)
		if err != nil {
			ErrorLogger.Printf("database error %v", err)
			return
		}
		_, err = db.Exec(createTable)
		if err != nil {
			ErrorLogger.Printf("database error %v", err)
			return
		}
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

	data, err := db.Query("SELECT usage FROM CPU_USAGE WHERE time >= UNIXEPOCH('now','-" + strconv.Itoa(seconds) + " seconds');")
	if err != nil {
		ErrorLogger.Printf("Cannot extract usage data from database: %v", err)
	}

	defer data.Close()
	var cpuUsages []float64
	for data.Next() {
		var usage string
		err = data.Scan(&usage)
		if err != nil {
			ErrorLogger.Printf("Cannot scan database: %v\n", err)
		}

		usage = strings.Replace(usage, "[", "", -1)
		usage = strings.Replace(usage, "]", "", -1)
		usages := strings.Split(usage, ",")

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

func withExecTime(hf httpHandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer duration(time.Now(), r.Method, r.URL)
		hf(w, r)
	})
}
func duration(start time.Time, method string, url *url.URL) {
	elapsed := time.Since(start)

	InfoLogger.Printf("%-5s | %-12s | %-20s ", method, elapsed, url)
}
