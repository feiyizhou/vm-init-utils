package linux_services

import (
	"fmt"
	"vm-init-utils/common"
	"vm-init-utils/options"
	"vm-init-utils/utils"
)

type LinuxService struct{}

func NewLinuxService() *LinuxService {
	return &LinuxService{}
}

func (ls *LinuxService) ReSetNetWork(network *options.Network) {
	switch network.OsType {
	case common.Centos:
		NewCentosService().ResetNetwork(network)
	default:
		utils.DieWithMsg(true, fmt.Sprintf("Unsupported os type: %s", network.OsType))
	}
	return
}

func (ls *LinuxService) SetPWD(sys *options.Sys) error {
	return nil
}

func (ls *LinuxService) SetHostname(hostname string) error {
	return utils.ExecCmd("hostnamectl", []string{"set-hostname", hostname})
}
