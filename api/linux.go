package api

type LinuxApi interface {
	SetHostname(hostname string) error
}
