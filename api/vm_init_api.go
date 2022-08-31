package api

type VMInitApi interface {
	SetIP() error
	SetPWD() error
}
