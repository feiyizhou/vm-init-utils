package api

type Network struct {
	Name     string `json:"name"`
	Source   string `json:"source"`
	Addr     string `json:"addr"`
	Mask     string `json:"mask"`
	Gateway  string `json:"gateway"`
	Gwmetric string `json:"gwmetric"`
}

type Sys struct {
	UserName string `json:"userName"`
	PWD      string `json:"pwd"`
}

type VMInitApi interface {
	SetIP() error
	SetPWD() error
}
