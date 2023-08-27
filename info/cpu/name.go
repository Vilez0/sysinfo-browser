package cpu

import (
	"bufio"
	"os"
	"strings"

	"github.com/Edip1/sysinfo-browser/util"
)

// Name returns the model name of the CPU.
func Name() string {
	// Open the file.
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		// Log the error.
		util.ErrorLogger.Printf("Error when opening file: %v\n", err)
	}
	// Scan the file.
	scanner := bufio.NewScanner(file)
	defer file.Close()

	// Scan the file line by line.
	for scanner.Scan() {
		// Check if the line starts with "model name: ".
		if strings.HasPrefix(scanner.Text(), "model name	: ") {
			// Return the text after "model name: ".
			return scanner.Text()[13:]
		}
	}
	return ``
}
