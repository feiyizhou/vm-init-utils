package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"runtime"
	"vm-init-utils/api"
	"vm-init-utils/common"
)

var ipChangeCmd = &cobra.Command{
	Use:   "set-ip",
	Short: "set-ip",
	Long:  "set-ip",
	Run: func(cmd *cobra.Command, args []string) {
		switch runtime.GOOS {
		case common.LINUX:
			err := setLinuxIP()
			if err != nil {
				log.Fatalf("Set linux ip err, err : %v \n", err)
				return
			}
		case common.WINDOWS:
			err := setWindowsIP()
			if err != nil {
				log.Fatalf("Set windows ip err, err : %v \n", err)
				return
			}
		default:
			log.Fatalln("Unknown os kind")
		}
	},
}

func setLinuxIP() error {
	linux := &api.Linux{}
	return linux.SetIP()
}

func setWindowsIP() error {
	windows := &api.Windows{
		Network: api.Network{
			Name:     "",
			Source:   "",
			Addr:     "",
			Mask:     "",
			Gateway:  "",
			Gwmetric: "",
		},
	}
	return windows.SetIP()
}
