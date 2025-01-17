package wrapper

import (
	"os"
	"os/exec"
)

func sendSignal(pid int, signal os.Signal) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	return process.Signal(signal)
}

func setProcessGroupID(cmd *exec.Cmd) {
	cmd.SysProcAttr.Setpgid = true
}
