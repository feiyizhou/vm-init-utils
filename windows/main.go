package main

import (
	"vm-init-utils/windows/cmd"
	"vm-init-utils/windows/utils"
)

func main() {
	utils.DoOrDieWithMsg(cmd.NewWinRootCmd().Execute(), "VM set failed")
}
