package util

import (
	"bufio"
	"os"
	"strings"
)

func ReadProcFile(filename string, info string) []string {
	file, err := os.Open(filename)
	if err != nil {
		ErrorLogger.Printf("Error when opening file: %v\n", err)
	}
	var lines []string

	scanner := bufio.NewScanner(file)
	defer file.Close()

	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), info) {
			lines = append(lines, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		ErrorLogger.Printf("Error when scanning file:%v\n", err)
	}
	return lines
}
