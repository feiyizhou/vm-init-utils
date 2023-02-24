package applications

import (
	"fmt"
	"vm-init-utils/common"
	"vm-init-utils/linux_services"
	"vm-init-utils/options"
	"vm-init-utils/utils"
	"vm-init-utils/windows_services"
)

var networkStrategyMap = map[string]ResetNetworkStrategy{
	common.LINUX:   linux_services.NewLinuxService(),
	common.WINDOWS: windows_services.NewWindowsService(),
}

type ResetNetworkStrategy interface {
	ReSetNetWork(conf *options.Network)
}

type ResetNetworkPipeline struct {
	strategy ResetNetworkStrategy
}

func NewResetNetworkPipeline(os string) *ResetNetworkPipeline {
	var (
		strategy ResetNetworkStrategy
		ok       bool
	)
	if strategy, ok = networkStrategyMap[os]; !ok || strategy == nil {
		utils.DieWithMsg(true, fmt.Sprintf("Unknown runtime goos: %s", os))
	}
	return &ResetNetworkPipeline{
		strategy: strategy,
	}
}

func (rp *ResetNetworkPipeline) ResetNetwork(conf *options.Network) {
	rp.strategy.ReSetNetWork(conf)
}
