package cmd

import (
	"fmt"
	"github.com/google/martian/log"
	"github.com/spf13/cobra"
	"vm-init-utils/linux_services"
	"vm-init-utils/modules"
)

var (
	name    string
	macAddr string
	ipAddr  string
	mask    string
	gateway string
	dns     string
)

func init() {
	ipChangeCmd.Flags().StringVarP(&name, "name", "n", "", "The name of interface")
	ipChangeCmd.Flags().StringVarP(&macAddr, "macAddr", "m", "", "The mac address of interface")
	ipChangeCmd.Flags().StringVarP(&ipAddr, "ipAddr", "i", "", "The ipv4 address of interface")
	ipChangeCmd.Flags().StringVarP(&mask, "mask", "s", "", "The mask of interface")
	ipChangeCmd.Flags().StringVarP(&gateway, "gateway", "g", "", "The gateway of interface")
	ipChangeCmd.Flags().StringVarP(&dns, "dns", "d", "", "The dns of interface, eg: 192.168.168.1,192.168.168.2")
}

var ipChangeCmd = &cobra.Command{
	Use:   "set-ip",
	Short: "set-ip",
	Long:  "set-ip",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(ipAddr) == 0 {
			return fmt.Errorf("Must notify a valid ipv4 address ")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		network := &modules.Network{
			Name:    name,
			MACAddr: macAddr,
			IPAddr:  ipAddr,
			NETMask: mask,
			GateWay: gateway,
			DNSStr:  dns,
		}
		err := linux_services.NewLinuxService().SetNetWork(network)
		if err != nil {
			log.Errorf("Set linux ip err, err : %v \n", err)
			return
		}
	},
}
