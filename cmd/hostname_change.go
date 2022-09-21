package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"runtime"
	"vm-init-utils/common"
	"vm-init-utils/config"
	"vm-init-utils/services"
	"vm-init-utils/utils"
)

var hostnameChangeCmd = &cobra.Command{
	Use:   "set-hostname",
	Short: "set-hostname",
	Long:  "set-hostname",
	Run: func(cmd *cobra.Command, args []string) {
		confFilePath := ""
		if len(args) != 0 {
			confFilePath = args[0]
		}
		hostname := config.GetSystemConf(confFilePath).Hostname
		switch runtime.GOOS {
		case common.LINUX:
			if len(hostname) == 0 {
				log.Println("Hostname is empty")
				return
			}
			err := services.NewLinuxService().SetHostname(hostname)
			utils.DoOrDieWithMsg(err, "Set hostname failed")
		default:
			utils.DieWithMsg(true, "Unsupported os type")
		}
	},
}
