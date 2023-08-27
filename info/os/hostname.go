package osutils

import (
	"os"

	"github.com/Edip1/sysinfo-browser/util"
)

func Hostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		util.ErrorLogger.Printf("cannot get hostname: %v\n", err)
	}
	return hostname
}
