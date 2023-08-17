package info

import (
	"bufio"
	"htop/util"
	"os"
	"strings"
)

func parseOsrelease() []string {
	var lines []string
	file, err := os.Open("/etc/os-release")
	if err != nil {
		file, err = os.Open("/usr/lib/os-release")
		if err != nil {
			util.ErrorLogger.Printf("Cannot open files /etc/os-release and /usr/lib/os-release: %v\n", err)
		}
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func OsName() string {
	osRelease := parseOsrelease()
	for e := range osRelease {
		if strings.Contains(osRelease[e], "NAME") {
			name := strings.Split(osRelease[e], "=")[1]
			name = strings.ReplaceAll(name, `"`, ``)
			return name
		}
	}
	return ""
}
