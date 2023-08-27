package disk

import (
	"bufio"
	"os"
	"strconv"

	"github.com/Edip1/sysinfo-browser/util"
)

// This code reads the size of the disk and converts it to GB.
func Size(disk string) int {
	var size int
	filepath := "/sys/block/" + disk + "/size"
	sizeFile, err := os.Open(filepath)
	if err != nil {
		util.ErrorLogger.Println("Error when reading file: ", filepath)
	}
	scanner := bufio.NewScanner(sizeFile)
	for scanner.Scan() {
		size, err = strconv.Atoi(scanner.Text())
		if err != nil {
			util.ErrorLogger.Println("Error when reading file: ", filepath)
		}
		size = size / 1953125
	}
	return size
}
