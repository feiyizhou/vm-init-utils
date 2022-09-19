package api

type Network struct {
	Name    string `json:"name"`
	Addr    string `json:"addr"`
	Mask    string `json:"mask"`
	Gateway string `json:"gateway"`
	DNS1    string `json:"dns1"`
	DNS2    string `json:"dns2"`
}

type Sys struct {
	UserName string `json:"userName"`
	PWD      string `json:"pwd"`
}

type VMInitApi interface {
	SetIP() error
	SetPWD() error
}
