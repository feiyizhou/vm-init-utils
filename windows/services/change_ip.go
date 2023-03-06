package services

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
	"strconv"
	"strings"
	"vm-init-utils/windows/options"
	"vm-init-utils/windows/utils"
)

type WindowsService struct{}

func NewWindowsService() *WindowsService {
	return &WindowsService{}
}

func (ws *WindowsService) ReSetNetWork(network *options.Network) error {
	var (
		ori, des   *options.Network
		err        error
		decodeType string
		gbkBytes   []byte
	)
	ori, err = gerOriNetworkConf()
	if err != nil {
		return err
	}
	des = coverConf(ori, network)

	// transform decode type of network interface's name
	if decodeType, err = utils.GetWinDecodeType(); err != nil {
		return err
	}
	if strings.EqualFold(decodeType, utils.GBKDecodeCode) {
		log.Infof("Local decode type is gbk, transform network interface name to gbk decode")
		if gbkBytes, err = utils.Utf8ToGbk([]byte(ori.Name)); err != nil {
			log.Errorf("transform network interface name to gbk decode failed, err: %v \n", err)
			return err
		}
		ori.Name = string(gbkBytes)
	}

	if err = executeResetNetworkIPAddrConf(des); err != nil {
		return err
	}
	return executeResetNetworkDNSConf(network)
}

func executeResetNetworkDNSConf(conf *options.Network) error {
	if len(conf.DNS) == 0 {
		return nil
	}
	var (
		err                                  error
		dhcpArgs, primaryArgs, assistantArgs []string
		dnsArr                               = strings.Split(conf.DNS, ",")
	)
	if len(dnsArr) == 0 {
		log.Infof("The dns config is nil, will not change the original config")
		return nil
	}
	dhcpArgs = []string{"interface", "ip", "set", "dns", fmt.Sprintf("\"%s\"", conf.Name), "dhcp"}
	if err = utils.ExecCmd("netsh", dhcpArgs); err != nil {
		return utils.MadeErr(err, "Reset dns server to dhcp failed")
	}
	primaryArgs = []string{
		"interface", "ip", "set", "dns",
		fmt.Sprintf("\"%s\"", conf.Name), "static", dnsArr[0], "primary",
	}
	if err = utils.ExecCmd("netsh", primaryArgs); err != nil {
		return utils.MadeErr(err, "Reset dns primary dns server failed")
	}
	if len(dnsArr[1:]) > 0 {
		for _, dns := range dnsArr[1:] {
			assistantArgs = []string{
				"interface", "ip", "add", "dns",
				fmt.Sprintf("\"%s\"", conf.Name), dns,
			}
			if err = utils.ExecCmd("netsh", assistantArgs); err != nil {
				return utils.MadeErr(err, fmt.Sprintf("Add assistant dns server failed, dns server: %s", dns))
			}
		}
	}
	return err
}

func executeResetNetworkIPAddrConf(conf *options.Network) error {
	var err error
	args := []string{
		"interface", "ip", "set", "address",
		fmt.Sprintf("\"%s\"", conf.Name),
		"source=static",
		fmt.Sprintf("addr=%s", conf.IPAddr),
		fmt.Sprintf("mask=%s", conf.NETMask),
		fmt.Sprintf("gateway=%s", conf.GateWay),
	}
	log.Infof("Reset network conf args: %v", args)
	if err = utils.ExecCmd("netsh", args); err != nil {
		err = utils.MadeErr(err, "Failed reset network conf")
	}
	return err
}

func coverConf(ori, new *options.Network) *options.Network {
	if len(new.Name) == 0 {
		new.Name = ori.Name
	}
	if len(new.NETMask) == 0 {
		new.NETMask = ori.NETMask
	}
	if len(new.GateWay) == 0 {
		new.GateWay = ori.GateWay
	}
	return new
}

func gerOriNetworkConf() (*options.Network, error) {
	var (
		inters     []net.Interface
		err        error
		addresses  []net.Addr
		ori        = &options.Network{}
		maskNumStr string
		maskStrArr []string
	)
	inters, err = net.Interfaces()
	if err != nil {
		return nil, utils.MadeErr(err, "Failed to get all network interfaces")
	}
	for _, inter := range inters {
		addresses, err = inter.Addrs()
		if err != nil {
			log.Errorf(fmt.Sprintf("Failed to get all address of %s", inter.Name))
			continue
		}
		for _, address := range addresses {
			networkIP, ok := address.(*net.IPNet)
			if ok && !networkIP.IP.IsLoopback() && networkIP.IP.To4() != nil {
				ori.Name = inter.Name
				ori.IPAddr = networkIP.IP.To4().String()
				maskNumStr = networkIP.Mask.String()
				if len(maskNumStr) != 8 {
					continue
				}
				for i := 0; i < 4; i++ {
					content, err := strconv.ParseUint(maskNumStr[i*2:(i+1)*2], 16, 32)
					if err != nil {
						log.Errorf("Parse mask num string to string failed")
						break
					}
					maskStrArr = append(maskStrArr, strconv.Itoa(int(content)))
				}
				if len(maskStrArr) == 4 {
					ori.NETMask = strings.Join(maskStrArr, ".")
				}
			}
		}
	}
	gateway, err := discoverGatewayOSSpecific()
	if err != nil {
		log.Errorf("Failed to get gateway")
	} else {
		ori.GateWay = gateway.To4().String()
	}
	if len(ori.IPAddr) == 0 || len(ori.NETMask) == 0 || len(ori.GateWay) == 0 {
		return nil, utils.MadeErr(nil, "Get ori network config failed")
	}
	return ori, err
}

func discoverGatewayOSSpecific() (ip net.IP, err error) {
	var output []byte
	output, err = utils.ExecCMDWithResult("route", []string{"print", "0.0.0.0"}, 10)
	utils.DoOrDieWithMsg(err, "Execute the command of get gateway failed")
	return parseWindowsGatewayIP(output)
}

func parseWindowsGatewayIP(output []byte) (net.IP, error) {
	parsedOutput, err := parseToWindowsRouteStruct(output)
	utils.DoOrDieWithMsg(err, "Failed to Parse output to windows route struct")
	ip := net.ParseIP(parsedOutput.Gateway)
	utils.DoOrDieWithMsg(err, "Failed to parse ip string to ip struct")
	return ip, nil
}

func parseToWindowsRouteStruct(output []byte) (windowsRouteStruct, error) {
	// Windows route output format is always like this:
	// ===========================================================================
	// Interface List
	// 8 ...00 12 3f a7 17 ba ...... Intel(R) PRO/100 VE Network Connection
	// 1 ........................... Software Loopback Interface 1
	// ===========================================================================
	// IPv4 Route Table
	// ===========================================================================
	// Active Routes:
	// Network Destination        Netmask          Gateway       Interface  Metric
	//           0.0.0.0          0.0.0.0      192.168.1.1    192.168.1.100     20
	// ===========================================================================
	//
	// Windows commands are localized, so we can't just look for "Active Routes:" string
	// I'm trying to pick the active route,
	// then jump 2 lines and get the row
	// Not using regex because output is quite standard from Windows XP to 8 (NEEDS TESTING)
	lines := strings.Split(string(output), "\n")
	sep := 0
	for idx, line := range lines {
		if sep == 3 {
			// We just entered the 2nd section containing "Active Routes:"
			if len(lines) <= idx+2 {
				return windowsRouteStruct{}, errNoGateway
			}

			fields := strings.Fields(lines[idx+2])
			if len(fields) < 5 {
				return windowsRouteStruct{}, errCantParse
			}

			return windowsRouteStruct{
				Destination: fields[0],
				Netmask:     fields[1],
				Gateway:     fields[2],
				Interface:   fields[3],
				Metric:      fields[4],
			}, nil
		}
		if strings.HasPrefix(line, "=======") {
			sep++
			continue
		}
	}
	return windowsRouteStruct{}, errNoGateway
}

type windowsRouteStruct struct {
	Destination string
	Netmask     string
	Gateway     string
	Interface   string
	Metric      string
}

var (
	errNoGateway = errors.New("no gateway found")
	errCantParse = errors.New("can't parse string output")
)
