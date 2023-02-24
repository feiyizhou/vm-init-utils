package modules

type Network struct {
	Name    string   `json:"name"`
	MACAddr string   `json:"macAddr"`
	IPAddr  string   `json:"ipAddr"`
	NETMask string   `json:"netMask"`
	GateWay string   `json:"gateWay"`
	UUID    string   `json:"uuid"`
	DNS     []string `json:"dns"`
	DNSStr  string   `json:"dnsStr"`
}

type Sys struct {
	UserName string `json:"userName"`
	PWD      string `json:"pwd"`
}
