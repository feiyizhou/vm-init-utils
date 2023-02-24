package common

const (
	LINUX   = "linux"
	WINDOWS = "windows"
)

const (
	Centos                = "centos"
	Ubuntu                = "ubuntu"
	OSTypeFlagFilePath    = "/etc/redhat-release"
	CentosNetConfFilePath = "/etc/sysconfig/network-scripts/ifcfg-"
)

type FlagSetName string

const (
	ResetNetworkFlagSet FlagSetName = "resetMNetwork"
)
