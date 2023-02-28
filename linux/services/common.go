package services

import (
	"fmt"
	"vm-init-utils/linux/common"
	"vm-init-utils/linux/options"
	"vm-init-utils/linux/utils"
)

type LinuxService struct{}

func NewLinuxService() *LinuxService {
	return &LinuxService{}
}

func (ls *LinuxService) ReSetNetWork(network *options.Network) error {
	switch network.OsType {
	case common.Centos:
		return NewCentosService().ResetNetwork(network)
	default:
		return utils.MadeErr(nil, fmt.Sprintf("Unsupported os type: %s", network.OsType))
	}
}

func (ls *LinuxService) SetPWD(sys *options.Sys) error {
	return nil
}

func (ls *LinuxService) SetHostname(hostname string) error {
	return utils.ExecCmd("hostnamectl", []string{"set-hostname", hostname})
}
