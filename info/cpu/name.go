package cpu

import (
	"bufio"
	"htop/util"
	"os"
	"strings"
)

func Name() string {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		util.ErrorLogger.Printf("Error when opening file: %v\n", err)
	}
	scanner := bufio.NewScanner(file)
	defer file.Close()

	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "model name	: ") {
			return scanner.Text()[13:]
		}
	}
	return ``
}
