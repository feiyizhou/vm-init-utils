package options

import "github.com/spf13/pflag"

type Network struct {
	OsType  string `json:"osType"`
	Name    string `json:"name" windows:"required"`
	MACAddr string `json:"macAddr"`
	IPAddr  string `json:"ipAddr" centos:"required" windows:"required"`
	NETMask string `json:"netMask"`
	GateWay string `json:"gateWay"`
	UUID    string `json:"uuid"`
	DNS     string `json:"dns"`
	DHCP    string `json:"dhcp"`
}

func NewNetworkFlags() *Network {
	return &Network{}
}

func (f *Network) AddFlags(mainfs *pflag.FlagSet) {
	fs := pflag.NewFlagSet("", pflag.ExitOnError)
	defer func() {
		fs.VisitAll(func(f *pflag.Flag) {
			if len(f.Deprecated) > 0 {
				f.Hidden = false
			}
		})
		mainfs.AddFlagSet(fs)
	}()
	fs.StringVarP(&f.Name, "name", "n", "", "The name of network interface")
	fs.StringVarP(&f.MACAddr, "macAddr", "m", "", "The mac address of network interface")
	fs.StringVarP(&f.IPAddr, "ipAddr", "i", "", "The ipv4 address of network interface")
	fs.StringVarP(&f.NETMask, "netmask", "k", "", "The netmask address of network interface")
	fs.StringVarP(&f.GateWay, "gateway", "g", "", "The gateway address of network interface")
	fs.StringVarP(&f.UUID, "uuid", "u", "", "The uuid of network interface, change this value supported on linux only")
	fs.StringVarP(&f.DHCP, "dhcp", "p", "false", "DNS address is from dhcp server or not, default false, must notify the dns value")
	fs.StringVarP(&f.DNS, "dns", "d", "", "The dns of destination machine, eg: 192.168.252.3,192.168.252.4")
}

type Sys struct {
	UserName string `json:"userName"`
	PWD      string `json:"pwd"`
}
