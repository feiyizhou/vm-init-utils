package main

import (
	"vm-init-utils/linux/cmd"
	"vm-init-utils/linux/utils"
)

func main() {
	utils.DoOrDieWithMsg(cmd.NewLinuxRootCmd().Execute(), "VM set failed")
}
