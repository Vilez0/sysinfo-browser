package disk

import (
	"bufio"
	"htop/util"
	"os"
	"strconv"
)

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
			size= size/1953125
		}
	return size
}
