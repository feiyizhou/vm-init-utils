package main

import "vm-init-utils/windows_services"

func main() {
	//err := cmd.NewChangeCmd().Execute()
	//if err != nil {
	//	log.Fatalf("Os init failed, err : %v \n", err)
	//}
	_, _ = windows_services.GetMACAddress()
}
