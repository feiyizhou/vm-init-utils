package utils

import (
	"os"
	"os/exec"
	"vm-init-utils/common"
)

func ExecShell(commands string, args []string, in, out, erf *os.File) error {
	command := exec.Command(commands, args...)
	if in == nil {
		in = os.Stdin
	}
	if out == nil {
		out = os.Stdout
	}
	if erf == nil {
		erf = os.Stderr
	}
	command.Stdin = in
	command.Stdout = out
	command.Stderr = erf
	return command.Run()
}

func GetOSType() (string, error) {
	_, err := os.Stat(common.OSTypeFlagFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return common.Ubuntu, nil
		}
		return "", err
	}
	return common.Centos, nil
}
