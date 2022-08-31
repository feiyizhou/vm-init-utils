package api

import "github.com/vishvananda/netlink"

type Linux struct {
	Network Network `json:"network"`
	Sys     Sys     `json:"sys"`
}

func (l *Linux) SetIP() error {
	link, err := netlink.LinkByName(l.Network.Name)
	if err != nil {
		return err
	}
	addr, err := netlink.ParseAddr(l.Network.Addr)
	if err != nil {
		return err
	}
	return netlink.AddrAdd(link, addr)
}

func (l *Linux) SetPWD() error {
	return nil
}
