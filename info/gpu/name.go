package gpu

import (
	"os/exec"
	"strings"

	"github.com/Edip1/sysinfo-browser/util"
)

func Name() string {
	var gpus string
	cmd, err := exec.Command("lspci").Output()
	if err != nil {
		util.ErrorLogger.Println("Error: ", err)
	}
	out := strings.Split(string(cmd), "\n")
	for _, e := range out {
		if strings.Contains(e, "VGA") {
			e := e[:len(e)-9][35:] + "\n"
			gpus += e
		}
	}

	return gpus
}
