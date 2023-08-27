package disk

import (
	"os"
	"strings"

	"github.com/Edip1/sysinfo-browser/util"
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
