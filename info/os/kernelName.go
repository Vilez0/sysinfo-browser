package info

import (
	"os/exec"
)

func KernelName() string {
	output, _ := exec.Command("uname", "-r").Output()
	return string(output)
}
