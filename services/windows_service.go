package services

import "vm-init-utils/modules"

type WindowsService struct{}

func NewWindowsService() *WindowsService {
	return &WindowsService{}
}

func (ws *WindowsService) SetNetWork(conf *modules.Network) error {
	return nil
}

func (ws *WindowsService) SetPWD(conf *modules.Sys) error {
	return nil
}

//func (w *modules.Windows) SetNetWork() error {
//	ipArgs := []string{
//		"interface",
//		"ip",
//		"set",
//		"address",
//		w.Network.Name,
//		"static",
//		w.Network.Addr,
//		w.Network.Mask,
//		w.Network.Gateway,
//	}
//	ipCmd := exec.Command("netsh", ipArgs...)
//	ipCmd.Stdin = os.Stdin
//	ipCmd.Stdout = os.Stdout
//	ipCmd.Stderr = os.Stderr
//	err := ipCmd.Run()
//	if err != nil {
//		return err
//	}
//	dnsArgs := []string{
//		"interface",
//		"ip",
//		"set",
//		"dns",
//		w.Network.Name,
//		"static",
//		w.Network.Gateway,
//	}
//	dnsCmd := exec.Command("netsh", dnsArgs...)
//	dnsCmd.Stdin = os.Stdin
//	dnsCmd.Stdout = os.Stdout
//	dnsCmd.Stderr = os.Stderr
//	return dnsCmd.Run()
//}
//
//func (w *Windows) SetPWD() error {
//	args := []string{
//		"user",
//		w.Sys.UserName,
//		w.Sys.PWD,
//	}
//	cmd := exec.Command("net", args...)
//	cmd.Stdin = os.Stdin
//	cmd.Stdout = os.Stdout
//	cmd.Stderr = os.Stderr
//	return cmd.Run()
//}
