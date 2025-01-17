package wrapper

import (
	"os"
	"os/exec"
	"syscall"
)

func sendSignal(pid int, _ os.Signal) error {
	// Sending signal is not supported on Windows. Instead, we will send a CTRL-C event.
	d, err := syscall.LoadDLL("kernel32.dll")
	if err != nil {
		return err
	}
	p, err := d.FindProc("GenerateConsoleCtrlEvent")
	if err != nil {
		return err
	}
	r, _, err := p.Call(syscall.CTRL_BREAK_EVENT, uintptr(pid))
	if r == 0 {
		return err
	}
	return nil
}

func setProcessGroupID(cmd *exec.Cmd) {
	cmd.SysProcAttr.CreationFlags = syscall.CREATE_NEW_PROCESS_GROUP
}
