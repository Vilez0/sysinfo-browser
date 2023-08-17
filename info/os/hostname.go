package osutils

import (
	"htop/util"
	"os"
)

func Hostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		util.ErrorLogger.Printf("cannot get hostname: %v\n",err)
	}
	return hostname
}
