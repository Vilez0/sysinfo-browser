package disk

import (
	"htop/util"
	"os"
	"strings"
)

func Disks() []string {
	entries, err := os.ReadDir("/sys/block/")
	if err != nil {
		util.ErrorLogger.Fatal(err)
	}
	var names []string
	for _, e := range entries {
		if !strings.Contains(e.Name(), "zram") {
			names = append(names, e.Name())
		}
	}
	return names
}
