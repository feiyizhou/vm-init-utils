package api

import "vm-init-utils/modules"

type CommonApi interface {
	// 根据网卡名称设置虚拟机的网卡信息，包括ip、netmask、mac、gateway、dns
	SetNetWork(conf *modules.Network) error
	SetPWD(conf *modules.Sys) error
}
