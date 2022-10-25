package linux_services

import "vm-init-utils/utils"

func (ls *LinuxService) SetHostname(hostname string) error {
	args := []string{
		"set-hostname",
		hostname,
	}
	return utils.ExecShell("hostnamectl", args, nil, nil, nil)
}
