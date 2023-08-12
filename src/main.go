package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shirou/gopsutil/cpu"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

type httpHandlerFunc func(http.ResponseWriter, *http.Request)

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
	http.Handle("/", withExecTime(getIndex))
	http.Handle("/index.mjs", withExecTime(getIndexjs))
	http.Handle("/index.css", withExecTime(getIndexCss))
	http.Handle("/realtime/cpus", withExecTime(cpusGet))
	InfoLogger.Println("starting server")
	err := http.ListenAndServe(":7052", nil)
	if err != nil {
		println(err)
	}
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func getIndexCss(w http.ResponseWriter, _ *http.Request) {
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

func getIndexjs(w http.ResponseWriter, _ *http.Request) {
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
func cpusGet(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "%s\n", getCpuUsage())
}

func storeCpuUsage(cpuUsageArray []byte) {
	const createTable string = `
  CREATE TABLE IF NOT EXISTS CPU_USAGE (
  time INTEGER NOT NULL PRIMARY KEY,
  usage BLOB NOT NULL
  );`
	cpuUsage := fmt.Sprintf("%s", cpuUsageArray)
	db, err := sql.Open("sqlite3", "../db/usage.db")
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
