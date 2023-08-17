package gpu

import (
	"htop/util"
	"os/exec"
	"strings"
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
			//00:02.0 VGA compatible controller: Intel Corporation 3rd Gen Core processor Graphics Controller (rev 09)
			e := e[:len(e)-9][35:] + "\n"
			gpus += e
		}
	}

	return gpus
	// return t1 + " " + t2
}
