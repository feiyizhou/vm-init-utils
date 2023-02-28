package cmd

import (
	"github.com/spf13/cobra"
	"runtime"
	"vm-init-utils/linux/common"
	"vm-init-utils/linux/utils"
)

var hostnameChangeCmd = &cobra.Command{
	Use:   "set-hostname",
	Short: "set-hostname",
	Long:  "set-hostname",
	Run: func(cmd *cobra.Command, args []string) {
		switch runtime.GOOS {
		case common.LINUX:
			utils.DieWithMsg(true, "Set hostname failed")
		default:
			utils.DieWithMsg(true, "Unsupported os type")
		}
	},
}
