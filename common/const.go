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
NAME="%s"
UUID="%s"
DEVICE="%s"
ONBOOT="yes"
MACADDR=%s
IPADDR=%s
NETMASK=%s
GATEWAY=%s
DNS1=%s
DNS2=%s
`
)

const (
	YamlConfigHomePath = "."
	YamlConfigName     = "config"
	YamlConfigType     = "yaml"
	YamlSysConfigKey   = "system"

	NetshOutputFilePath = "c://netsh_output"
)
