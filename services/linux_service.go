package services

import (
	"bufio"
	"fmt"
	"github.com/vishvananda/netlink"
	"io"
	"os"
	"strings"
	"vm-init-utils/common"
	"vm-init-utils/modules"
	"vm-init-utils/utils"
)

type LinuxService struct{}

func NewLinuxService() *LinuxService {
	return &LinuxService{}
}

func (ls *LinuxService) SetHostname(hostname string) error {
	args := []string{
		"set-hostname",
		hostname,
	}
	return utils.ExecShell("hostnamectl", args, nil, nil, nil)
}

func (ls *LinuxService) SetNetWork(network *modules.Network) error {
	osType, err := utils.GetOSType()
	utils.DoOrDieWithMsg(err, "get os type info error")
	switch osType {
	case common.Centos:
		var filePath string
		if len(network.Name) == 0 {
			linkName, err := getLinkName()
			if err != nil {
				return err
			}
			if len(linkName) == 0 {
				return fmt.Errorf("No network interface founded ")
			}
			filePath = fmt.Sprintf("%s%s", common.CentosNetConfFilePath, linkName)
		} else {
			filePath = fmt.Sprintf("%s%s", common.CentosNetConfFilePath, network.Name)
		}

		// 获取原始配置
		oriConf, err := getOriNetworkConf(filePath)
		utils.DieWithMsg(err != nil, "Get centos original network config error")
		utils.DieWithMsg(oriConf == nil, "Centos original network config is empty")

		// 删除原始文件
		err = os.Remove(filePath)
		utils.DieWithMsg(err != nil, "Delete the origin conf file failed")

		// 覆盖配置
		newConf := rewriteConf(oriConf, network)

		// 创建新文件并写入新的网络配置
		confStr := fmt.Sprintf(common.CentosNetConfTemplate, newConf.Name, newConf.UUID, newConf.Name,
			newConf.MACAddr, newConf.IPAddr, newConf.NETMask, newConf.GateWay, newConf.DNS1,
			newConf.DNS2)
		newFile, err := os.Create(filePath)
		utils.DoOrDieWithMsg(err, "Create new config file failed")
		_, err = newFile.WriteString(confStr)
		utils.DoOrDieWithMsg(err, "Write new network config to file failed")
		defer newFile.Close()

		// 重启网络服务
		args := []string{
			"restart",
			"network",
		}
		err = utils.ExecShell("systemctl", args, nil, nil, nil)
		utils.DieWithMsg(err != nil, "Restart network service failed")
	case common.Ubuntu:
		utils.DieWithMsg(true, "Unsupported os type")
	default:
		utils.DieWithMsg(strings.EqualFold("", osType), "Unknown os type")
	}
	return nil
}

func (ls *LinuxService) SetPWD(sys *modules.Sys) error {
	return nil
}

func rewriteConf(oriConf, network *modules.Network) *modules.Network {
	if len(network.MACAddr) != 0 {
		oriConf.MACAddr = network.MACAddr
	}
	if len(network.IPAddr) != 0 {
		oriConf.IPAddr = network.IPAddr
	}
	if len(network.NETMask) != 0 {
		oriConf.NETMask = network.NETMask
	}
	if len(network.GateWay) != 0 {
		oriConf.GateWay = network.GateWay
	}
	if len(network.DNS1) != 0 {
		oriConf.DNS1 = network.DNS1
	} else if len(oriConf.DNS) != 0 {
		oriConf.DNS1 = oriConf.DNS
	}
	if len(network.DNS2) != 0 {
		oriConf.DNS2 = network.DNS2
	}
	return oriConf
}

func getOriNetworkConf(filePath string) (*modules.Network, error) {
	oriFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	if oriFile == nil {
		return nil, fmt.Errorf("Network interface config file is not exist ")
	}
	defer oriFile.Close()

	oriConf := &modules.Network{}
	br := bufio.NewReader(oriFile)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if len(a) == 0 {
			continue
		}
		tmpStr := strings.Split(string(a), "=")[1]
		if strings.Contains(string(a), "NAME") {
			oriConf.Name = strings.ReplaceAll(tmpStr, "\"", "")
		}
		if strings.Contains(string(a), "UUID") {
			oriConf.UUID = strings.ReplaceAll(tmpStr, "\"", "")
		}
		if strings.Contains(string(a), "MACADDR") {
			oriConf.MACAddr = tmpStr
		}
		if strings.Contains(string(a), "IPADDR") {
			oriConf.IPAddr = tmpStr
		}
		if strings.Contains(string(a), "NETMASK") {
			oriConf.NETMask = tmpStr
		}
		if strings.Contains(string(a), "GATEWAY") {
			oriConf.GateWay = tmpStr
		}
		if strings.Contains(string(a), "DNS") {
			oriConf.DNS = tmpStr
		}
		if strings.Contains(string(a), "DNS1") {
			oriConf.DNS1 = tmpStr
		}
		if strings.Contains(string(a), "DNS2") {
			oriConf.DNS2 = tmpStr
		}
	}
	return oriConf, err
}

func getLinkName() (string, error) {
	var name string
	linkArr, err := netlink.LinkList()
	if err != nil {
		return "", err
	}
	for _, link := range linkArr {
		attrs := link.Attrs()
		if strings.EqualFold(attrs.EncapType, "ether") {
			name = attrs.Name
			break
		}
	}
	return name, err
}
