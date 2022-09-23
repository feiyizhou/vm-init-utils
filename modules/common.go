package modules

type Network struct {
	Name    string `json:"name"`
	MACAddr string `json:"macAddr"`
	IPAddr  string `json:"ipAddr"`
	NETMask string `json:"netMask"`
	GateWay string `json:"gateWay"`
	UUID    string `json:"uuid"`
	DNS     string `json:"dns"`  // 保留字段
	DNS1    string `json:"dns1"` // 主DNS
	DNS2    string `json:"dns2"` // 副DNS
}

type Sys struct {
	UserName string `json:"userName"`
	PWD      string `json:"pwd"`
}
