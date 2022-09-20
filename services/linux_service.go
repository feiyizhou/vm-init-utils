package services

import (
	"bufio"
	"fmt"
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

func (ls *LinuxService) SetNetWork(network *modules.Network) error {
	osType, err := utils.GetOSType()
	utils.DoOrDieWithMsg(err, "get os type info error")
	var uuid string
	switch osType {
	case common.Centos:
		filePath := fmt.Sprintf("%s%s", common.CentosNetConfFilePath, network.Name)
		// 读取文件
		oriFile, err := os.Open(filePath)
		utils.DieWithMsg(err != nil, "Open centos network config file error")
		utils.DieWithMsg(oriFile == nil, "Centos network config file is empty")
		defer oriFile.Close()
		br := bufio.NewReader(oriFile)
		for {
			a, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}
			if strings.Contains(string(a), "UUID") {
				uuid = strings.Split(string(a), "=")[1]
			}
		}

		// 备份原始文件
		err = os.Rename(filePath, fmt.Sprintf("%s-origin", filePath))
		utils.DoOrDieWithMsg(err, "Rename network config file failed")

		// 创建新文件并写入新的网络配置
		confStr := fmt.Sprintf(common.CentosNetConfTemplate, network.Name, uuid, network.Name,
			network.MACAddr, network.IPAddr, network.NETMask, network.GateWay, network.DNS1,
			network.DNS2)
		newFile, err := os.Create(filePath)
		utils.DoOrDieWithMsg(err, "Create new config file failed")
		_, err = newFile.WriteString(confStr)
		utils.DoOrDieWithMsg(err, "Write new network config to file failed")
		defer newFile.Close()

		// 重启网络服务
		args := []string{
			"network",
			"restart",
		}
		err = utils.ExecShell("service", args, nil, nil, nil)
		utils.DieWithMsg(err != nil, "Restart network service failed")
	case common.Ubuntu:

	default:
		utils.DieWithMsg(strings.EqualFold("", osType), "Unknown os type")
	}
	return nil
}

func (ls *LinuxService) SetPWD(sys *modules.Sys) error {
	return nil
}
