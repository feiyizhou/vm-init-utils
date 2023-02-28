package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"runtime"
	"vm-init-utils/linux/common"
)

var pwdChangeCmd = &cobra.Command{
	Use:   "set-pwd",
	Short: "set-pwd",
	Long:  "set-pwd",
	Run: func(cmd *cobra.Command, args []string) {
		switch runtime.GOOS {
		case common.LINUX:
			err := setLinuxPwd()
			if err != nil {
				log.Fatalf("Set linux pwd err, err : %v \n", err)
				return
			}
		case common.WINDOWS:
			err := setWindowsPwd()
			if err != nil {
				log.Fatalf("Set windows pwd err, err : %v \n", err)
				return
			}
		default:
			log.Fatalln("Unknown os kind")
		}
	},
}

func setLinuxPwd() error {
	return nil
}

func setWindowsPwd() error {
	return nil
}
