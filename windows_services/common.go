package windows_services

import "vm-init-utils/options"

type WindowsService struct{}

func NewWindowsService() *WindowsService {
	return &WindowsService{}
}

func (ws *WindowsService) ReSetNetWork(network *options.Network) {

}

func getNetworkConf() (*options.Network, error) {

	return nil, nil
}
