package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"runtime"
	"vm-init-utils/common"
	"vm-init-utils/config"
	"vm-init-utils/linux_services"
	"vm-init-utils/modules"
	"vm-init-utils/windows_services"
)

var ipChangeCmd = &cobra.Command{
	Use:   "set-ip",
	Short: "set-ip",
	Long:  "set-ip",
	Run: func(cmd *cobra.Command, args []string) {
		confFilePath := ""
		if len(args) != 0 {
			confFilePath = args[0]
		}
		conf := config.GetSystemConf(confFilePath).Network
		network := &modules.Network{
			Name:    conf.Name,
			MACAddr: conf.MACAddr,
			IPAddr:  conf.IPAddr,
			NETMask: conf.NETMASK,
			GateWay: conf.GATEWAY,
			DNS1:    conf.DNS1,
			DNS2:    conf.DNS2,
		}
		switch runtime.GOOS {
		case common.LINUX:
			err := linux_services.NewLinuxService().SetNetWork(network)
			if err != nil {
				log.Fatalf("Set linux ip err, err : %v \n", err)
				return
			}
		case common.WINDOWS:
			err := windows_services.NewWindowsService().SetNetWork(network)
			if err != nil {
				log.Fatalf("Set windows ip err, err : %v \n", err)
				return
			}
		default:
			log.Fatalln("Unknown os kind")
		}
	},
}
