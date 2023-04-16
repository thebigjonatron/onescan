package scans

import "net"

type Scanner interface {
	Start(utils UtilsArp)
}

type DefaultScan struct {
}

func (*DefaultScan) Start(utils UtilsArp) {
	arp := new(ArpScan)
	syn := new(SynScan)
	arp.Start(utils)
	// Need to modify utils
	syn.Start(utils)
	print(arp, syn)
	//Run default scan. Arp scan and syn scan
}

type UtilsArp struct {
	Intface *net.Interface
	Ports   []int
}
type UtilsSyn struct {
}
