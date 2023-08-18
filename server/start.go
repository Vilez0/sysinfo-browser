package server

import (
	"encoding/json"
	"htop/util"
	"reflect"

	"github.com/gin-gonic/gin"
)

func Run(addr string) {
	router := gin.Default()
	router.GET("/", ServeIndex)
	router.GET("/:file", ServeIndex)
	router.GET("/realtime/cpus/", ServeCpuUsage)
	router.GET("/realtime/cpus/:seconds/*average", ServeCpuUsage)
	router.GET("/system/:name/:info", ServeSystem)
	router.Run(addr)
}

func marshaler(value any) string {

	rt := reflect.TypeOf(value)
	if rt.Kind() != reflect.Slice && rt.Kind() != reflect.Array {
		var slice []any
		slice = append(slice, value)
		value = slice
	}
	jsonValue, err := json.Marshal(value)
	if err != nil {
		util.ErrorLogger.Println(`cannot marshal json: `, err)
	}
	return string(jsonValue)
}
