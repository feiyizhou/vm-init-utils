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

const (
	CentosNetConfTemplate = `
TYPE="Ethernet"
PROXY_METHOD="none"
BROWSER_ONLY="no"
BOOTPROTO="static"
DEFROUTE="yes"
IPV4_FAILURE_FATAL="no"
IPV6INIT="yes"
IPV6_AUTOCONF="yes"
IPV6_DEFROUTE="yes"
IPV6_FAILURE_FATAL="no"
IPV6_ADDR_GEN_MODE="stable-privacy"
ONBOOT="yes"
NAME="%s"
UUID="%s"
DEVICE="%s"
IPADDR=%s
NETMASK=%s
GATEWAY=%s

`
)
