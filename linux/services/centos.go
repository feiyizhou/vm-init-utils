package services

import (
	"fmt"
	"github.com/vishvananda/netlink"
	"os"
	"strings"
	"vm-init-utils/linux/common"
	"vm-init-utils/linux/options"
	"vm-init-utils/linux/utils"
)

type CentosService struct{}

func NewCentosService() *CentosService {
	return &CentosService{}
}

func (cs *CentosService) ResetNetwork(network *options.Network) error {
	var (
		filePath string
		err      error
		oriConf  map[string]string
		content  = strings.Builder{}
	)
	filePath, err = cs.getDefaultFilePath(network.Name)
	if err != nil {
		return utils.MadeErr(err, "Get default network interface config file error")
	}

	// 获取原始配置
	oriConf, err = cs.makeUpConf(filePath)
	if err != nil || oriConf == nil {
		return utils.MadeErr(err, "Get centos original network config error")
	}

	// 删除原始文件
	if err = os.Remove(filePath); err != nil {
		return utils.MadeErr(err, "Delete the origin conf file failed")
	}
	if utils.IsExist(fmt.Sprintf("%s-bck", filePath)) {
		if err = os.Remove(fmt.Sprintf("%s-bck", filePath)); err != nil {
			return utils.MadeErr(err, "Delete the origin conf back up file failed")
		}
	}

	// 重写配置
	if !strings.Contains(oriConf["BOOTPROTO"], "static") {
		content.WriteString(fmt.Sprintf(common.CentosNetConfTemplate, oriConf["NAME"], oriConf["UUID"], oriConf["DEVICE"],
			network.IPAddr, network.NETMask, network.GateWay))
		if len(network.MACAddr) > 0 {
			content.WriteString(fmt.Sprintf("HWADDR=%s\n", network.MACAddr))
		}
		if len(network.DNS) > 0 {
			dnsArr := strings.Split(network.DNS, ",")
			content.WriteString(fmt.Sprintf("DNS=%s\n", dnsArr[0]))
			for index, dns := range dnsArr[1:] {
				content.WriteString(fmt.Sprintf("DNS%d=%s\n", index+1, dns))
			}
		}
		var file *os.File
		defer file.Close()
		if file, err = os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0666); err != nil {
			return err
		}
		if _, err = file.WriteString(content.String()); err != nil {
			return err
		}
	} else {
		if err = cs.rewriteNewConf(cs.overrideConf(oriConf, network), filePath); err != nil {
			return utils.MadeErr(err, "Rewrite network configuration failed")
		}
	}

	// 重启网络服务
	if err = utils.ExecCmd("systemctl", []string{"restart", "network"}); err != nil {
		return utils.MadeErr(err, "Restart network service failed")
	}
	return nil
}

func (cs *CentosService) getDefaultFilePath(name string) (string, error) {
	var (
		filePath, defaultName string
		err                   error
	)
	defaultName, err = cs.getLinkName()
	if err != nil {
		return "", err
	}
	if len(name) == 0 {
		if len(defaultName) == 0 {
			return "", fmt.Errorf("No network interface founded ")
		}
		filePath = fmt.Sprintf("%s%s", common.CentosNetConfFilePath, defaultName)
	} else {
		if len(defaultName) == 0 {
			return "", fmt.Errorf("No network interface founded ")
		}
		if err = checkDefaultLinKMAC(defaultName); err != nil {
			return "", err
		}
		filePath = fmt.Sprintf("%s%s", common.CentosNetConfFilePath, name)
	}
	return filePath, nil
}

func checkDefaultLinKMAC(name string) error {
	var (
		filePath    = fmt.Sprintf("%s%s", common.CentosNetConfFilePath, name)
		err         error
		macConfAddr string
		content     []byte
		lines       []string
	)
	content, err = utils.ExecCMDWithResult("bash",
		[]string{"-c", fmt.Sprintf("ifconfig %s | grep ether | awk '{print $2}'", name)}, 10)
	if err != nil {
		return utils.MadeErr(err, "Get mac address failed")
	}
	lines, err = utils.ReadFileToLines(filePath)
	if err != nil {
		return utils.MadeErr(err, "Read default ether config file failed")
	}
	for _, line := range lines {
		if strings.Contains(line, "HWADDR") {
			macConfAddr = strings.Split(line, "=")[1]
			break
		}
	}
	if !strings.Contains(strings.ReplaceAll(string(content), " ", ""), macConfAddr) {
		err = utils.ExecCmd("bash", []string{"-c",
			fmt.Sprintf("sed -i /HWADDR/d %s", filePath),
		})
		if err != nil {
			return utils.MadeErr(err, "Delete default ether hardware addr config failed")
		}
	}
	return nil
}

func (cs *CentosService) rewriteNewConf(conf map[string]string, filePath string) error {
	var (
		file    *os.File
		err     error
		content = strings.Builder{}
	)
	defer file.Close()
	for k, v := range conf {
		if strings.Contains(k, "DNS") {
			for index, dns := range strings.Split(v, ",") {
				if index == 0 {
					content.WriteString(fmt.Sprintf("DNS=%s\n", dns))
				} else {
					content.WriteString(fmt.Sprintf("DNS%d=%s\n", index, dns))
				}
			}
		} else if strings.Contains(k, "BOOTPROTO") {
			content.WriteString(fmt.Sprintf("%s=static\n", k))
		} else {
			content.WriteString(fmt.Sprintf("%s=%s\n", k, v))
		}
	}
	if file, err = os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0666); err != nil {
		return err
	}
	_, err = file.WriteString(content.String())
	return err
}

func (cs *CentosService) overrideConf(oriConfMap map[string]string, network *options.Network) map[string]string {
	for k, _ := range oriConfMap {
		switch true {
		case (strings.Contains(k, "MACADDR") || strings.Contains(k, "HWADDR")) && len(network.MACAddr) != 0:
			oriConfMap[k] = network.MACAddr
		case strings.Contains(k, "IPADDR") && len(network.IPAddr) != 0:
			oriConfMap[k] = network.IPAddr
		case strings.Contains(k, "NETMASK") && len(network.NETMask) != 0:
			oriConfMap[k] = network.NETMask
		case strings.Contains(k, "GATEWAY") && len(network.GateWay) != 0:
			oriConfMap[k] = network.GateWay
		case strings.Contains(k, "DNS") && len(network.DNS) != 0:
			oriConfMap[k] = network.DNS
		}
	}
	return oriConfMap
}

func (cs *CentosService) makeUpConf(filePath string) (map[string]string, error) {
	var (
		oriConfMap, oriConfBckMap map[string]string
		err                       error
		bckFilePath               = fmt.Sprintf("%s-bck", filePath)
	)
	if oriConfMap, err = cs.getOriNetworkConf(filePath); err != nil {
		return nil, fmt.Errorf("Failed to get original network config, err: %v ", err)
	}
	if utils.IsExist(bckFilePath) {
		if oriConfBckMap, err = cs.getOriNetworkConf(bckFilePath); err != nil {
			return nil, fmt.Errorf("Failed to get original network back up config, err: %v ", err)
		}
		for k, v := range oriConfBckMap {
			if _, ok := oriConfMap[k]; !ok {
				oriConfMap[k] = v
			}
		}
	}
	return oriConfMap, err
}

func (cs *CentosService) getOriNetworkConf(filePath string) (map[string]string, error) {
	var (
		lines []string
		err   error
	)
	if !utils.IsExist(filePath) {
		return nil, nil
	}
	if lines, err = utils.ReadFileToLines(filePath); err != nil {
		return nil, err
	}
	return cs.parseConfLines(lines), nil
}

func (cs *CentosService) parseConfLines(lines []string) map[string]string {
	var (
		dnsArr []string
		conf   = make(map[string]string)
	)
	for _, line := range lines {
		line = strings.ReplaceAll(line, "\n", "")
		line = strings.ReplaceAll(line, " ", "")
		if len(line) == 0 {
			continue
		}
		if !strings.Contains(line, "=") {
			continue
		}
		kvArr := strings.Split(line, "=")
		k, v := kvArr[0], kvArr[1]
		if !strings.Contains(k, "DNS") {
			conf[k] = v
		} else {
			dnsArr = append(dnsArr, v)
		}
	}
	if len(dnsArr) != 0 {
		conf["DNS"] = strings.Join(dnsArr, ",")
	}
	return conf
}

func (cs *CentosService) getLinkName() (string, error) {
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
