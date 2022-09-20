package modules

type Network struct {
	Name    string `json:"name"`
	MACAddr string `json:"macAddr"`
	IPAddr  string `json:"ipAddr"`
	NETMask string `json:"netMask"`
	GateWay string `json:"gateWay"`
	DNS1    string `json:"dns1"`
	DNS2    string `json:"dns2"`
}

type Sys struct {
	UserName string `json:"userName"`
	PWD      string `json:"pwd"`
}
