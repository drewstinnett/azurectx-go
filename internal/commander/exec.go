package commander

import "os/exec"

type Commander interface {
	Output(string, ...string) ([]byte, error)
}

type RealCommander struct{}

func CaptureStdOut(cmd string, args ...string) (string, error) {
	out, err := exec.Command(cmd, args...).Output()
	return string(out), err
}

// mock cmd.Execute
func (c RealCommander) Output(command string, args ...string) ([]byte, error) {
	return exec.Command(command, args...).Output()
}
